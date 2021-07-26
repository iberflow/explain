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

	//args := os.Args[1:]

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

		_ = matchers.NewMatcher(cmd, parsedPage.Options)
		//argumentMatcher := matchers.NewMatcher(cmd, result.Options)
		//spew.Dump(argumentMatcher.Match())
		app := gui.NewApp(parsedPage.Options, cmd)
		app.Draw()
	}

}
