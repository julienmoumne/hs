package options

import (
	"bytes"
	"github.com/jessevdk/go-flags"
)

type OptionParser struct {
	options Options
	parser  *flags.Parser
}

func (p *OptionParser) Parse() (Options, error) {
	p.createOptionsParser()
	var _, err = p.parser.Parse()
	return p.options, err
}

func (p *OptionParser) createOptionsParser() {
	p.options = Options{}
	p.parser = flags.NewParser(&p.options, flags.HelpFlag)
	p.parser.Name = "hs"
	p.parser.Usage = "[-f <arg>...] [Options]"
	p.parser.ShortDescription = "Interactive single keystroke menus for the shell"
	p.parser.LongDescription = "Hotshell is a command-line application to efficiently recall and share commands."
}

func (p *OptionParser) CreateManPage() string {
	p.createOptionsParser()
	var buf bytes.Buffer
	p.parser.WriteManPage(&buf)
	return buf.String()
}
