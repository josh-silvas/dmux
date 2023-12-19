package main

import (
	"github.com/josh-silvas/nbot/core"
	"github.com/josh-silvas/nbot/plugins/info"
	"github.com/josh-silvas/nbot/plugins/keystore"
	sshinteractive "github.com/josh-silvas/nbot/plugins/ssh"
	"github.com/josh-silvas/nbot/plugins/upgrade"
	"github.com/josh-silvas/nbot/plugins/version"
	"github.com/josh-silvas/nbot/shared/keyring"
	"github.com/sirupsen/logrus"
)

var buildVersion = "0.0.1+dev"

func main() {

	// Create a new cli parser and register all the plugins to be used.
	// This is where the arg commands are defined and the func to execute
	// when called.
	parser := core.NewParser(
		info.Plugin,
		keystore.Plugin,
		sshinteractive.Plugin,
		upgrade.Plugin,
		version.Plugin,
	)

	// Get the keyring configuration file from the
	// default store location (homedir/.config/gokeys)
	cfg, err := keyring.New(logrus.Debug)
	if err != nil {
		logrus.Fatalf("nbot.keyring.New:%s", err)
	}
	cfg.Meta = map[string]string{
		"buildVersion": buildVersion,
	}

	// Run a check of the current version. This will only alert and perform
	// a check against artifactory every 2 hours.
	if err := version.Check(cfg); err != nil {
		logrus.Warning(err)
	}

	// Run the parser to parse all the arguments defined by nbot and
	// the additional plugins. This will also check if and what argument happened
	// and execute the defined plugin function.
	parser.Run(cfg)
}
