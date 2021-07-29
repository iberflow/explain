package text

import (
	"github.com/ignasbernotas/explain/config"
	"github.com/ignasbernotas/explain/parsers/args"
	"github.com/ignasbernotas/explain/parsers/man"
	"strings"
)

func hasArgument(arg string, list *man.List) bool {
	name := strings.TrimLeft(arg, "-")

	for _, opt := range list.Options() {
		// search for exact match
		if opt.Name == name || opt.Alias == name {
			return true
		}
	}

	return false
}

func formatArgument(regionIndex int, arg *args.Arg, list *man.List) string {
	// match double dashed args
	if strings.HasPrefix(arg.Name, "--") && hasArgument(arg.Name, list) {
		return MarkRegion(regionIndex, arg.Name, true)
	}

	// match single dash arg
	if strings.HasPrefix(arg.Name, "-") && hasArgument(arg.Name, list) {
		return MarkRegion(regionIndex, arg.Name, true)
	}

	// split single dash arg into multiple single-char args
	if strings.HasPrefix(arg.Name, "-") {
		name := strings.TrimLeft(arg.Name, "-")

		var subset []string
		// split arg into characters and search
		for ii, c := range name {
			cc := string(c)
			if hasArgument(cc, list) {
				subset = append(subset, MarkRegion(regionIndex+1*100+ii, cc, true))
			} else {
				subset = append(subset, cc)
			}
		}

		return "-" + strings.Join(subset, "")
	}

	return config.ColorArg(arg.Name, false, "")
}

func RenderCommand(cmd *args.Command, list *man.List) string {
	var found []string

	for i, arg := range cmd.Args.Items() {
		found = append(found, formatArgument(i, arg, list))
	}

	return cmd.Name + " " + strings.Join(found, " ")
}
