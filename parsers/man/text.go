package man

import (
	"regexp"
	"strings"
)

var space = regexp.MustCompile(`\s`)
var replacements = map[string]string{
	MacroSquareBracketStart: "[",
	MacroSquareBracketEnd:   "]",
	MacroNoSpacesOn:         "",
	MacroNoSpacesOff:        "",
	MacroArgument:           "",
	MacroArgListStart:       "",
	MacroArgListEnd:         "",
}

func replace(str string) string {
	for token, replacement := range replacements {
		str = strings.ReplaceAll(str, "."+token, replacement)
		str = strings.ReplaceAll(str, token, replacement)
	}

	return str
}

func isArgumentList(str string) bool {
	return strings.Contains(str, MacroArgListStart)
}

// .It Fl R Xo
// .Sm off
// .Oo Ar bind_address : Oc
// .Ar port : host : hostport
// .Sm on
// .Xc

// [bind_address:]port:host:hostport
func updateDescriptionAndName(opt *Option, str string) *Option {
	sep := "\n"
	argList := isArgumentList(opt.Name)
	noSpacesIndex := -1
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
		} else {
			l = replace(l)
			newLines = append(newLines, l)
		}
	}

	opt.Description = strings.Join(newLines, sep)

	return opt
}

func getIndicator(name string) string {
	if strings.Contains(name, "--") {
		return "--"
	}

	if strings.Contains(name, "-") {
		return "-"
	}

	return "-"
}
