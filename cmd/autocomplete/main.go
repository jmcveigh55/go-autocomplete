package main

import (
	"github.com/alecthomas/kong"

	"github.com/jmcveigh55/autocomplete/internal/autocomplete"
)

var cli struct {
	File        string `arg:"" help:"The file containing all possible completions." type:"existingfile"`
	PartialWord string `arg:"" help:"The query word fragment to be completed."`
}

func main() {
	kong.Parse(&cli)
	autocomplete.Run(cli.File, cli.PartialWord)
}
