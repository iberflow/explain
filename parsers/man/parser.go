package man

import (
	"regexp"
	"strings"
)

type Page struct {
	Name    string
	Options *OptionList
}

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

var sectionPattern = regexp.MustCompile(`(?i)\.` + MacroStructureSectionHeading + `\s(.*)`)

// curl doc uses .IP for options
// ssh doc uses .It fl for options
var optionPattern = regexp.MustCompile(`(?i)\.I[P|t]\s(fl)?(.*)`)

func (p *Parser) Parse(str string) *Page {
	var currentSectionName string
	var currentArgName string
	var currentArgDescription string
	var innerStructure bool

	page := &Page{
		Options: &OptionList{},
	}

	for _, line := range strings.Split(str, "\n") {
		line = strings.TrimSpace(line)

		// start of inner content
		// we can ignore everything within it
		// there are probably more of these inner content sections that we need to ignore
		if isMacro(line, MacroStructureRelativeInsetStart) {
			innerStructure = true
			continue
		}

		// end of inner content
		if isMacro(line, MacroStructureRelativeInsetEnd) {
			innerStructure = false
			continue
		}

		if p.isSectionLine(line) {
			currentSectionName = p.parseSectionLine(line)
		}

		// only look for arguments in specific sections
		if (currentSectionName == SectionDescription || currentSectionName == SectionOptions) && p.isOptionLine(line) {
			arg := p.parseOptionLine(line)

			// if this is not the very first arg in the doc
			// store the previous one
			// also ignore args without descriptions, they're not helpful
			if len(currentArgName) > 0 {
				if len(currentArgDescription) > 0 {
					opt := NewOption(currentArgName, currentArgDescription)
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
	return optionPattern.MatchString(str)
}

func (p *Parser) parseOptionLine(str string) string {
	r := optionPattern.FindStringSubmatch(str)
	if len(r) == 0 {
		return ""
	}

	value := strings.Trim(r[2], `" `)

	return value
}
