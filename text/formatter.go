package text

import (
	"fmt"
	"github.com/ignasbernotas/explain/config"
	"github.com/ignasbernotas/explain/parsers/args"
	"github.com/ignasbernotas/explain/parsers/man"
	"regexp"
	"strings"
)

func FormatDescription(desc string) string {
	var pattern = regexp.MustCompile(`(?i)(?m)\\f` + man.MacroFontItalic + `([^\\fP]+.*?)\\fP`)
	matches := pattern.FindAllStringSubmatch(desc, -1)
	for i, m := range matches {
		if args.IsArg(m[1]) {
			desc = strings.Replace(desc, m[0], MarkRegion(i, m[1], true), 1)
		}
	}

	return desc
}

var colorRemovePattern = regexp.MustCompile(`\[.*\](.*?)\[.*\]`)

func StripColor(desc string) string {
	matches := colorRemovePattern.FindAllStringSubmatch(desc, -1)
	for _, m := range matches {
		desc = strings.Replace(desc, m[0], m[1], 1)
	}

	return desc
}

func ColorOption(index int, opt *man.Option) string {
	return Underline(MarkRegion(index, opt.String(), true)) + "[::d]" + opt.StringArg()
}

func Underline(arg string) string {
	return fmt.Sprintf(`[::u]%s[""]`, arg)
}

func MarkRegion(index int, arg string, flag bool) string {
	return fmt.Sprintf(`["%d"]%s[""]`, index, config.ColorArg(arg, flag))
}
