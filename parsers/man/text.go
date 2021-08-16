package man

import (
	"regexp"
	"strings"
)

var space = regexp.MustCompile(`\s`)
var spaceMulti = regexp.MustCompile(`\s+`)
var wrapReplacements = map[string]string{
	MacroFontUnderline:  `[::u]$2[::-]`,
	MacroArgWithoutDash: `[::b]$2[::-]`,
	MacroBrackets:       `($2)`,
	MacroFlag:           `[::b]-$2[::-]`,
	MacroDoubleQuote:    `"$2"`,
	MacroSingleQuote:    `'$2'`,
}

var spacedReplacements = map[string]string{
	MacroArgument:       "",
	MacroFontUnderline:  "",
	MacroArgWithoutDash: "",
	MacroUnix:           "UNIX",
	MacroNoSpace:        "",
	//MacroDoubleQuoteOpen:  `"`,
	//MacroDoubleQuoteClose: `"`,
}

var replacements = map[string]string{
	MacroSquareBracketStart: "[",
	MacroSquareBracketEnd:   "]",
	MacroNoSpacesOn:         "",
	MacroNoSpacesOff:        "",
	MacroArgListStart:       "",
	MacroArgListEnd:         "",
	MacroBrackets:           "",
	MacroParagraph2:         "\n",
	MacroDoubleQuote:        "",
	MacroSingleQuote:        "",
	MacroUnix + " " + MacroNoSpace + " ": "UNIX",
}

var wrapRef = map[string]string{
	MacroManPageReference: "$2($3)", //.Xr syslog 3
	// .Xr gzip 1 .
}

func wrapReference(str string) string {
	for token, replacement := range wrapRef {
		var pat2 = regexp.MustCompile(`(?i)(\.?` + token + `)\s([\w|\-_]+)\s([\w|\-_]+)`)
		str = pat2.ReplaceAllString(str, replacement)
	}
	return str
}

func replaceToolName(str, toolName string) string {
	var pat = regexp.MustCompile(`(?i)(\.?` + MacroName + `)`)
	str = pat.ReplaceAllString(str, " "+toolName)

	return str
}

func wrapReplace(str string) string {
	for token, replacement := range wrapReplacements {
		var pat = regexp.MustCompile(`(?i)(\.?` + token + `)\s(.*)`)
		str = pat.ReplaceAllString(str, replacement)
	}

	return str
}

func replaceTokens(str string) string {
	for token, replacement := range replacements {
		str = strings.ReplaceAll(str, "."+token, replacement)
		str = strings.ReplaceAll(str, token, replacement)
		str = strings.ReplaceAll(str, token+"\n", replacement+"\n")
		str = strings.ReplaceAll(str, token+" ", replacement)
	}

	for token, replacement := range spacedReplacements {
		str = strings.ReplaceAll(str, "."+token, replacement)
		str = strings.ReplaceAll(str, token+" ", replacement)
		str = strings.ReplaceAll(str, " "+token, replacement)
	}

	return str
}

func replace(str string) string {
	str = wrapReference(str)
	str = wrapReplace(str)
	str = replaceTokens(str)

	return str
}

func isArgumentList(str string) bool {
	return strings.Contains(str, MacroArgListStart)
}

func updateAttributes(opt *Option, str, toolName string) *Option {
	sep := "\n"
	noSpacesIndex := -1
	argList := isArgumentList(opt.Name)
	if argList {
		opt.Name = strings.ReplaceAll(opt.Name, " "+MacroArgListStart, "")
		opt.Alias = strings.ReplaceAll(opt.Alias, " "+MacroArgListStart, "")
	}
	lines := strings.Split(str, sep)
	var newLines []string
	for index, l := range lines {
		if l == MacroArgListStart {
			argList = true
		}
		if l == MacroArgListEnd {
			argList = false
		}
		if l == MacroNoSpacesOn {
			noSpacesIndex = index
		}

		if l == MacroNoSpacesOff {
			if noSpacesIndex < 0 {
				noSpacesIndex = 0
			}

			lineWithNoSpaces := space.ReplaceAllString(strings.Join(newLines[noSpacesIndex:index], ""), "")
			lineWithNoSpaces = replace(lineWithNoSpaces)

			newLines = newLines[:noSpacesIndex-1]

			if argList {
				if len(opt.Parameters) == 0 {
					opt.Parameters = []string{lineWithNoSpaces}
				}
			} else {
				newLines = append(newLines, lineWithNoSpaces)
			}
			continue
		}

		newLines = append(newLines, replace(l))
	}

	opt.Description = strings.TrimSpace(strings.Join(newLines, sep))
	opt.Description = fixSentences(opt.Description)
	opt.Description = replaceToolName(opt.Description, toolName)
	opt.Description = spaceMulti.ReplaceAllString(opt.Description, " ")

	return opt
}

func fixSentences(str string) string {
	str = strings.ReplaceAll(str, " ,", ",")
	str = strings.ReplaceAll(str, " .", ".")
	str = strings.ReplaceAll(str, " ( ", " (")
	str = strings.ReplaceAll(str, `  `, ` `)

	// TODO: tview does not display \n\n in text
	var pattern = regexp.MustCompile(`(?i)(\.?\n)+`)
	str = pattern.ReplaceAllStringFunc(str, func(s string) string {
		if strings.Contains(s,".\n") {
			return s
		}

		return " "
	})
	return str
}
