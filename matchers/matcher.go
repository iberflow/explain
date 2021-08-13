package matchers

import (
	"github.com/ignasbernotas/explain/parsers/args"
	"github.com/ignasbernotas/explain/parsers/man"
	"strings"
)

type Matcher struct {
	command              *args.Command
	documentationOptions *man.List
}

func NewMatcher(command *args.Command, options *man.List) *Matcher {
	return &Matcher{
		command:              command,
		documentationOptions: options,
	}
}

func (m *Matcher) Match() *man.List {
	found := *man.NewList(make([]*man.Option, 0))

	for _, arg := range m.command.Args.Items() {
		// match double dashed args
		if strings.HasPrefix(arg.Name, "--") {
			name := strings.TrimLeft(arg.Name, "-")

			for _, opt := range m.documentationOptions.Options() {
				// search for exact match
				if opt.Name == name || opt.Alias == name {
					found.Add(opt)
					break
				}
			}

			break
		}

		// match single dash args
		if strings.HasPrefix(arg.Name, "-") {
			name := strings.TrimLeft(arg.Name, "-")

			for _, opt := range m.documentationOptions.Options() {
				// search for exact match
				if opt.Name == name || opt.Alias == name {
					found.Add(opt)
					break
				} else {
					// split arg into characters and search
					for _, c := range name {
						cc := string(c)
						if opt.Name == cc || opt.Alias == cc {
							found.Add(opt)
							break
						}
					}
				}
			}
		}
	}

	return found.Unique(false)
}
