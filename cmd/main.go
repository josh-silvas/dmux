package main

import (
	"github.com/josh-silvas/nbot/internal/core"
	"github.com/josh-silvas/nbot/internal/keyring"
	"github.com/josh-silvas/nbot/internal/nlog"
	"github.com/josh-silvas/nbot/plugins/info"
	"github.com/josh-silvas/nbot/plugins/keystore"
	sshinteractive "github.com/josh-silvas/nbot/plugins/ssh"
	"github.com/josh-silvas/nbot/plugins/upgrade"
	"github.com/josh-silvas/nbot/plugins/version"
)

var (
	buildVersion = "0.0.1+dev"
	l            = nlog.NewWithGroup(core.AppName)
)

func main() {

	// Create a new cli parser and register all the plugins to be used.
	// This is where the arg commands are defined and the func to execute
	// when called.
	parser := core.NewParser(
		new(info.Plugin),
		new(keystore.Plugin),
		new(sshinteractive.Plugin),
		new(upgrade.Plugin),
		new(version.Plugin),
	)

	// Get the keyring configuration file from the
	// default store location (homedir/.config/nbot)
	cfg, err := keyring.New(l)
	if err != nil {
		l.Fatalf("nbot.keyring.New:%s", err)
	}
	cfg.Meta = map[string]string{
		"buildVersion": buildVersion,
	}

	// Run a check of the current version. This will only alert and perform
	// a check against artifactory every 2 hours.
	if err := version.Check(cfg); err != nil {
		l.Warn(err)
	}

	// Run the parser to parse all the arguments defined by nbot and
	// the additional plugins. This will also check if and what argument happened
	// and execute the defined plugin function.
	parser.Run(cfg)
}
