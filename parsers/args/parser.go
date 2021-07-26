package args

import "strings"

const ArgColor = "#ff0000"
const FlagColor = "#fff000"

type Command struct {
	Name string
	Args []string
}

func NewCommand(str string) *Command {
	args := strings.Split(str, " ")

	return &Command{
		Name: args[0],
		Args: args[1:],
	}
}

func (c *Command) String() string {
	return c.Name + strings.Join(c.Args, " ")
}

func (c *Command) StringRegions() string {
	cmd := `[cmd]` + c.Name + `[""]`

	for _, arg := range c.Args {
		cmd += ` ` + c.ColorArg(arg)
	}

	return cmd
}

func (c *Command) ColorArg(arg string) string {
	color := ArgColor

	if strings.Contains(arg, "-") {
		color = FlagColor
	}

	arg = `[` + color + `]` + arg + `[""]`

	return arg
}

func Parse(str string) (commands []*Command) {
	pipedCommands := strings.Split(str, "|")

	for i := 0; i < len(pipedCommands); i++ {
		commands = append(commands, NewCommand(pipedCommands[i]))
	}

	return commands
}
