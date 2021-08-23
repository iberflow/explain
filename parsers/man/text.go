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
	MacroQuoteLiteral:   `"$2"`,
	MacroSingleQuote:    `'$2'`,
	MacroBeginList:      "",
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
	MacroSquareBracketStart:              "[",
	MacroSquareBracketEnd:                "]",
	MacroNoSpacesOn:                      "",
	MacroNoSpacesOff:                     "",
	MacroParagraphTag:                    "",
	MacroArgListStart:                    "",
	MacroArgListEnd:                      "",
	MacroBrackets:                        "",
	MacroParagraph2:                      "\n",
	MacroExtendedArgList:                 "\n",
	MacroEndList:                         "",
	MacroDoubleQuote:                     "",
	MacroSingleQuote:                     "",
	MacroUnix + " " + MacroNoSpace + " ": "UNIX",
	`\fB`:                                "",
	`\fI`:                                "",
	`\fR`:                                "",
	`\fP`:                                "",
}

var wrapRef = map[string]string{
	MacroManPageReference: "$2($3)", //.Xr syslog 3
	// .Xr gzip 1 .
}

func wrapReference(str string) string {
	for token, replacement := range wrapRef {
		var pattern = regexp.MustCompile(`(?i)(\.?` + token + `)\s([\w|\-_]+)\s([\w|\-_]+)`)
		str = pattern.ReplaceAllString(str, replacement)
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

func updateAttributes(opt *Option, str, toolName string) *Option {
	sep := "\n"
	noSpacesIndex := -1

	// some arguments have parameters that indicate a start of an arg list
	// reset those parameters and tell the parser to expect that list
	argList := false
	if len(opt.Parameters) > 0 && opt.Parameters[0] == MacroArgListStart {
		argList = true
		opt.Parameters = []string{}
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
	opt.Description = replace(opt.Description)

	opt.Name = replace(opt.Name)
	opt.Alias = replace(opt.Alias)

	return opt
}

func fixSentences(str string) string {
	str = strings.ReplaceAll(str, " ,", ",")
	str = strings.ReplaceAll(str, " .", ".")
	str = strings.ReplaceAll(str, " ( ", " (")
	str = strings.ReplaceAll(str, `  `, ` `)
	str = strings.ReplaceAll(str, `\&`, ``)

	// TODO: tview does not display multi-line text
	var pattern = regexp.MustCompile(`(?i)(\.?\n)+`)
	str = pattern.ReplaceAllStringFunc(str, func(s string) string {
		if strings.Contains(s, ".\n") {
			return s
		}

		return " "
	})

	return str
}
