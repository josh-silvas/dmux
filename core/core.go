// Package core contains the core functionality of nbot and initializes the default
// parser.
package core

import (
	"fmt"
	"os"
	"syscall"

	"github.com/akamensky/argparse"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/josh-silvas/nbot/core/keyring"
	"github.com/sirupsen/logrus"
)

const (
	// AppName is a constant for nbot so that we can be consistent on the
	// naming across all uses of the name.
	AppName = "nbot"
)

type (
	// Parser type is used as the main parser for nbot.
	// It embeds the argparse Parser type.
	Parser struct {
		*argparse.Parser
		Plugins []Plugin
	}
	// Plugin is the command and calling function for each plugin
	Plugin struct {
		CMD  *argparse.Command
		Func func(keyring.Settings)
	}
)

var debugFlag *bool

// NewParser function will initiate and return the parent parser for the
// nbot app.
func NewParser(fn ...func(*Parser) Plugin) Parser {
	// Create new main parser object
	p := Parser{
		Parser:  argparse.NewParser(AppName, "NBot ʘ‿ʘ: Nautobot CLI."),
		Plugins: make([]Plugin, 0),
	}

	// Define the top-level arguments pinned to the nbot parser.
	debugFlag = p.Flag("", "debug", &argparse.Options{Help: "view debug level logging"})

	// Register the plugin commands into the parser
	for _, f := range fn {
		p.Plugins = append(p.Plugins, f(&p))
	}
	return p
}

// Run method will parse the arguments in the parser as well as range through all the
// registered plugins to determine which action "Happened()"
func (p *Parser) Run(cfg keyring.Settings) {
	// Parse input
	if err := p.Parse(os.Args); err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(p.Usage(text.FgRed.Sprint(err)))
		// abort if there is an error parsing arguments
		syscall.Exit(1)
	}
	if *debugFlag {
		logrus.SetLevel(logrus.DebugLevel)
	}

	for _, v := range p.Plugins {
		if v.CMD.Happened() {
			v.Func(cfg)
		}
	}
}
