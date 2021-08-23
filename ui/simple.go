package ui

import (
	"github.com/pterm/pterm"
)

func SimpleUi(proc *Processor) {
	primary := pterm.NewStyle(pterm.FgYellow, pterm.BgBlack)

	primary.Println()
	primary.Println("Command: " + proc.command.String())
	primary.Println()

	for _, n := range proc.CommandOptions().Options() {
		pterm.DefaultSection.Println(n.StringWithArg())
		pterm.Description.Println(n.Description)
	}

	primary.Println()
}
