package main

import (
	"github.com/ignasbernotas/explain/gui"
	"github.com/ignasbernotas/explain/matchers"
	"github.com/ignasbernotas/explain/parsers/args"
	"github.com/ignasbernotas/explain/parsers/man"
	manreader "github.com/ignasbernotas/explain/readers/man"
	"log"
	"os"
)

func main() {
	var manPath = os.Getenv("MANPATH")

	str := "curl -sSL -a --basic 'https://install.larashed.com/linux'"

	commands := args.Parse(str)

	reader := manreader.NewReader(manPath)

	for _, cmd := range commands {
		manPage, err := reader.Read(cmd.Name)
		if err != nil {
			log.Println(err.Error())
			return
		}

		parser := man.NewParser()
		parsedPage := parser.Parse(manPage)
		argumentMatcher := matchers.NewMatcher(cmd, parsedPage.Options)

		app := gui.NewApp(parsedPage.Options, cmd, argumentMatcher.Match())
		app.Draw()
	}
}
