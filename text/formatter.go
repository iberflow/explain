package text

import (
	"fmt"
	"github.com/ignasbernotas/explain/config"
	"github.com/ignasbernotas/explain/parsers/man"
	"regexp"
	"strings"
)

func isArg(arg string) bool {
	return strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--")
}

func FormatDescription(desc string) string {
	var pattern = regexp.MustCompile(`(?i)(?m)\\f` + man.MacroFontItalic + `([^\\fP]+.*?)\\fP`)
	matches := pattern.FindAllStringSubmatch(desc, -1)
	for i, m := range matches {
		if isArg(m[1]) {
			desc = strings.Replace(desc, m[0], markRegion(i, m[1]), 1)
		}
	}

	return desc
}

var colorRemovePattern = regexp.MustCompile(`\[.*\](.*?)\[.*\]`)

func StripColor(desc string) string {
	matches := colorRemovePattern.FindAllStringSubmatch(desc, -1)
	for _, m := range matches {
		if isArg(m[1]) {
			desc = strings.Replace(desc, m[0], m[1], 1)
		}
	}

	return desc
}

func markRegion(index int, arg string) string {
	return fmt.Sprintf(`["%d"]%s[""]`, index, config.ColorArg(arg))
}
