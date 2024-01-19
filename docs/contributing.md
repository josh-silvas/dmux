# Contributing
Contributing to DMux is intended to be as easy as possible. If you have any idea on
how to make the development workflow better or easier, please submit a PR!

At a high-level, all development and testing is set up with helper `make` commands so that
the experience can be the same across developers. 

## Creating your own Plugin
DMux sub-commands are declared as plugins within the application. Each plugin should
have an exportable function named `Plugin` that accepts a `*core.Parser` argument
and returns a `core.Plugin` data type.

An example framework of a basic plugin would be:
```go
package cool_plugin

import (
	"fmt"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/josh-silvas/dmux/core"
)

// Arguments are globally defined in the plugin package, so that the
// invoker of your plugin can overwrite them with user-defined values.
var arg1 *string

func Plugin(p *core.Parser) core.Plugin {
	// This is what your main sub-command will be called.
	cmd := p.NewCommand("example", "An example plugin!")
	
	// Any additional arguments you would like to add.
	arg1 = cmd.String("a", "some-arg", &argparse.Options{Help: "Explains what some-arg does."})
	return core.Plugin{CMD: cmd, Func: pluginFunc}
}

func pluginFunc(cfg keyring.Settings) {
	// Within your `pluginFunc` is where all of the custom code will live 
	// for your plugin. 
	
	// When the command is invoked, this function will be called. 
	// NOTE: The `cfg keyring.Settings` argument that your pluginFunc will be passed,
	//       will give you the ability to access credentials stored in the DMux Keychain.
	
	// Example fetching the Nautobot api key from the DMux keychain:
	nKey, err := cfg.Nautobot()
}
```
> NOTE: All plugins are stored in the `plugins` directory.

Once your plugin is created, add it to the main parser in
[cmd/main.go](cmd/main.go), for example:
```go
	// Add your plugin to this list, alphabetically please!!
	parser := core.NewParser(
		ansible.Plugin,
		console.Plugin,
		keystore.Plugin,
		libre.Plugin,
		info.Plugin,
		sshinteractive.Plugin,
		version.Plugin,
	)
```

## Testing the Application 
All testing and linting for DMux are done in a development docker container so that we
may produce similar results across development environments.

When you run a `make` command, this repo will spin up the docker container as defined
in the `development/Dockerfile` with all the packages you need to run linting and unit-testing.

### Running Tests with Make
Each `make` command will run within the development container. If you would like
to run them outside of the development container. Run the command with a prepended 
underscore (`_`), for example, `make _lint`

| Command         | Description                                                |
|-----------------|------------------------------------------------------------|
| `make tests`    | Run all testing frameworks (same as CI pipeline).          |
| `make lint`     | Only run linting tests like golangci-lint and yamllint.    |
| `make unittest` | Only run unit-testing. Runs for entire repository.         |
| `make build`    | Rebuild the development container. If dependencies change. |
    
## Release a New Build
The `version` plugin in DMux will periodically make a call-out to Artifactory to 
determine if a newer release has been published. If so, the users will get a notification
to upgrade with the upgrade steps, depending on their architecture. 

To draft a new DMux release, you only need to publish the release on Github and Artifactory, and the 
users will update at their leisure. 

### Pre-Reqs for Releasing a New Build
Before you can release a new build, goreleaser needs credentials to the 
systems that it's going to communicate with to publish a new release. 

Please make sure you have the following defined

1. Github token with release permissions to [github.com/josh-silvas/dmux](https://github.com/josh-silvas/dmux).
   * Store the token in `~/.config/goreleaser/github_token`


### Steps for Releasing a New Build
1. Building of the Go binary 
2. Debian build for Linux installations
3. Brew build for OSX installations (this includes the formula.rb generation)
4. Archive build of tar.gz files
5. Drafting of a new release in [GitHub](https://github.com/josh-silvas/dmux/releases).

Once you have completed the prereqs, you should be able to draft a new release using make:
```
make release
```  
> Note: this requires a clean git status and a new git tag.  If you are wanting to test this release, run 
```
make test_release
```       


## More help
> Send an email to [Josh Silvas](mailto:josh@jsilvas.com)
