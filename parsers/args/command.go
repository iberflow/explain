package args

import (
	"strings"
)

type Arg struct {
	Name     string
	Partials []string
}

type Command struct {
	Name string
	Args *List
}

func NewCommand(str string) *Command {
	args := strings.Split(str, " ")

	name := args[0]
	args = args[1:]

	return &Command{
		Name: name,
		Args: parseArgs(args),
	}
}

func (c *Command) String() string {
	return c.Name + " " + c.Args.String()
}

func Parse(str string) (commands []*Command) {
	pipedCommands := strings.Split(str, "|")

	for i := 0; i < len(pipedCommands); i++ {
		commands = append(commands, NewCommand(strings.TrimSpace(pipedCommands[i])))
	}

	return commands
}

func parseArgs(args []string) *List {
	list := &List{items: []*Arg{}}

	for _, arg := range args {
		// some arguments may contain "chains" of flags e.g. curl -sSl
		a := &Arg{Name: arg, Partials: []string{}}

		if IsShortArg(arg) {
			var short []string

			arg = strings.TrimLeft(arg, "-")

			for i, _ := range arg {
				short = append(short, string(arg[i]))
			}

			a.Partials = short
		}

		list.Add(a)
	}

	return list
}
