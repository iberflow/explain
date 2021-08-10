package main

import (
	"github.com/ignasbernotas/explain/parsers/man"
	manreader "github.com/ignasbernotas/explain/readers/man"
	"github.com/ignasbernotas/explain/ui"
	"os"
	"strings"
)

func main() {
	reader := manreader.NewReader(os.Getenv("MANPATH"))

	args := os.Args[1:]
	cmd := strings.Join(args, " ")

	parser := man.NewParser()
	processor := ui.NewProcessor(reader, parser)
	processor.LoadCommand(cmd)

	app := ui.NewApp(processor)
	app.Draw()
}
