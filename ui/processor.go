package ui

import (
	"github.com/ignasbernotas/explain/matchers"
	"github.com/ignasbernotas/explain/parsers/args"
	"github.com/ignasbernotas/explain/parsers/man"
	manreader "github.com/ignasbernotas/explain/readers/man"
	"log"
)

type Processor struct {
	manPath string

	command        *args.Command
	commandOptions *man.List

	docOptions *man.List

	reader  *manreader.Reader
	parser  *man.Parser
	matcher *matchers.Matcher
}

func NewProcessor(reader *manreader.Reader, parser *man.Parser) *Processor {
	return &Processor{
		reader:     reader,
		parser:     parser,
		docOptions: &man.List{},
		command:    args.NewCommand(""),
	}
}

func (p *Processor) LoadCommand(command string) {
	p.command = args.Parse(command)

	manPage, err := p.reader.Read(p.command.Name)
	if err != nil {
		log.Println(err.Error())
		return
	}

	parsedPage := p.parser.Parse(manPage, p.command.Name)
	p.docOptions = parsedPage.Options
	p.docOptions = p.docOptions.Unique(true)
	p.docOptions.Sort()

	argumentMatcher := matchers.NewMatcher(p.command, p.docOptions)
	p.commandOptions = argumentMatcher.Match()
}

func (p *Processor) Command() *args.Command {
	return p.command
}

func (p *Processor) CommandOptions() *man.List {
	return p.commandOptions
}

func (p *Processor) DocumentationOptions() *man.List {
	return p.docOptions
}
