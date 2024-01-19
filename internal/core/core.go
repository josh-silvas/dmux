// Package core contains the core functionality of dmux and initializes the default
// parser.
package core

import (
	"fmt"
	"log/slog"
	"os"
	"syscall"

	"github.com/akamensky/argparse"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/josh-silvas/dmux/internal/keyring"
	"github.com/josh-silvas/dmux/internal/nlog"
)

const (
	// AppName is a constant for dmux so that we can be consistent on the
	// naming across all uses of the name.
	AppName = "dmux"
)

type (
	// Parser type is used as the main parser for dmux.
	// It embeds the argparse Parser type.
	Parser struct {
		*argparse.Parser
		Plugins []Plugin
	}

	// PluginBase : base struct for all plugins
	PluginBase struct {
		C   *argparse.Command
		Log nlog.Logger
	}

	// Plugin : interface for all plugins
	Plugin interface {
		Register(*Parser) Plugin
		CMD() *argparse.Command
		Func(keyring.Settings)
	}
)

var debugFlag *bool

// NewParser function will initiate and return the parent parser for the
// dmux app.
func NewParser(plugins ...Plugin) Parser {
	// Create new main parser object
	p := Parser{
		Parser:  argparse.NewParser(AppName, "DMux ʘ‿ʘ: Networking CLI."),
		Plugins: make([]Plugin, 0),
	}

	// Define the top-level arguments pinned to the dmux parser.
	debugFlag = p.Flag("", "debug", &argparse.Options{Help: "view debug level logging"})

	// Register the plugin commands into the parser
	for i := range plugins {
		p.Plugins = append(p.Plugins, plugins[i].Register(&p))
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
		nlog.LogLevel.Set(slog.LevelDebug)
	}

	for _, v := range p.Plugins {
		if v.CMD().Happened() {
			v.Func(cfg)
		}
	}
}
