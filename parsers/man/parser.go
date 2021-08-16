package man

import (
	"regexp"
	"strings"
)

type Page struct {
	Name    string
	Options *List
}

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

var sectionPattern = regexp.MustCompile(`(?i)\.` + MacroStructureSectionHeading + `\s(.*)`)

// curl doc uses .IP for options
// ssh doc uses .It fl for options
var curlPattern = regexp.MustCompile(`(?i)\.IP\s(.*)`)
var sshPattern = regexp.MustCompile(`(?i)\.It\sFl\s(.*)`)
var allOptionPattern = regexp.MustCompile(`(?i)\.I[P|t]\s(fl)?(.*)`)

const (
	TypeCurl = "curl"
	TypeSSH  = "ssh"
)

func (p *Parser) Parse(str, commandName string) *Page {
	var currentSectionName string
	var currentArgName string
	var currentArgDescription string
	var innerStructure bool

	page := &Page{
		Options: &List{},
	}

	for _, line := range strings.Split(str, "\n") {
		line = strings.TrimSpace(line)

		// start of inner content
		// we can ignore everything within it
		// there are probably more of these inner content sections that we need to ignore
		if isMacro(line, MacroStructureRelativeInsetStart) {
			innerStructure = true
			//fmt.Println("is macro start")
			continue
		}

		// end of inner content
		if isMacro(line, MacroStructureRelativeInsetEnd) {
			//fmt.Println("is macro end")

			innerStructure = false
			continue
		}

		if p.isSectionLine(line) {
			currentSectionName = p.parseSectionLine(line)
			//fmt.Println("name: " + currentSectionName)
		}

		// only look for arguments in specific sections
		if (currentSectionName == SectionDescription || currentSectionName == SectionOptions) && p.isOptionLine(line) {
			//fmt.Println("----------------------")
			//fmt.Println("line: ", line)
			arg := p.parseOptionLine(line)
			//fmt.Println("arg:", arg)

			// if this is not the very first arg in the doc
			// store the previous one
			// also ignore args without descriptions, they're not helpful
			if len(currentArgName) > 0 {
				if len(currentArgDescription) > 0 {
					//fmt.Println("opt: ", currentArgName, "desc: ", currentArgDescription)
					opt := NewOption(commandName, currentArgName, currentArgDescription)
					page.Options.Add(opt)
				}
			}

			currentArgName = arg
			currentArgDescription = "" // reset for next arg
		} else {
			if len(currentArgName) > 0 && !innerStructure {
				currentArgDescription += "\n" + line
			}
		}
	}

	return page
}

func (p *Parser) isSectionLine(str string) bool {
	return sectionPattern.MatchString(str)
}

func isMacro(text string, macro ...string) bool {
	var pattern = regexp.MustCompile(`(?i)\.(` + strings.Join(macro, "|") + `)(.*)`)

	return pattern.MatchString(text)
}

func (p *Parser) parseSectionLine(str string) string {
	r := sectionPattern.FindStringSubmatch(str)
	if len(r) == 0 {
		return ""
	}

	value := strings.TrimSpace(r[1])

	return value
}

func (p *Parser) isOptionLine(str string) bool {
	return allOptionPattern.MatchString(str)
}

func getPatternType(str string) string {
	if curlPattern.MatchString(str) {
		return TypeCurl
	}

	if sshPattern.MatchString(str) {
		return TypeSSH
	}

	return TypeCurl
}

func (p *Parser) parseOptionLine(str string) string {
	r := curlPattern.FindStringSubmatch(str)

	if len(r) == 0 {
		r = sshPattern.FindStringSubmatch(str)
		if len(r) == 0 {
			return ""
		}
		value := strings.Trim(r[1], `" `)
		return value
	}

	value := strings.Trim(r[2], `" `)

	return value
}
