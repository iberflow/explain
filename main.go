package main

import (
	"fmt"
	"github.com/ignasbernotas/explain/parsers/man"
	manreader "github.com/ignasbernotas/explain/readers/man"
	"github.com/ignasbernotas/explain/ui"
	"os"
	"strings"
)

func main() {
	reader := manreader.NewReader(os.Getenv("MANPATH"))
	simpleUI := true

	args := os.Args[1:]
	if args[0] == "-h" || args[0] == "--help" {
		fmt.Println("Use `explain -i <command>` for interactive mode.")
		return
	}

	if args[0] == "-i" {
		args = args[1:]
		simpleUI = false
	}

	cmd := strings.Join(args, " ")

	parser := man.NewParser()
	processor := ui.NewProcessor(reader, parser)
	err := processor.LoadCommand(cmd)

	if err != nil {
		fmt.Println(err.Error() + "\n" + "Use `explain -i <command>` for interactive mode.")
		return
	}

	if simpleUI {
		ui.SimpleUi(processor)
		return
	}

	app := ui.NewApp(processor)
	app.Draw()
}
