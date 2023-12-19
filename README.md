# NBot : Nautobot CLI Tool 
[![Go Report Card](https://goreportcard.com/badge/github.com/josh-silvas/nbot)](https://goreportcard.com/report/github.com/josh-silvas/nbot)
[![GoDoc](https://godoc.org/github.com/josh-silvas/nbot?status.svg)](https://godoc.org/github.com/josh-silvas/nbot)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/josh-silvas/nbot/blob/main/LICENSE)
[![Release](https://img.shields.io/github/release/josh-silvas/nbot.svg?style=flat-square)](https://github.com/josh-silvas/nbot/releases/latest)
[![Downloads](https://img.shields.io/github/downloads/josh-silvas/nbot/total.svg?style=flat-square)](https://github.com/josh-silvas/nbot/releases)
[![GitHub issues](https://img.shields.io/github/issues/josh-silvas/nbot.svg?style=flat-square)](https://github.com/josh-silvas/nbot/issues)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/josh-silvas/nbot.svg?style=flat-square)](https://github.com/josh-silvas/nbot/pulls)
[![GitHub contributors](https://img.shields.io/github/contributors/josh-silvas/nbot.svg?style=flat-square)](https://github.com/josh-silvas/nbot/graphs/contributors)
[![GitHub stars](https://img.shields.io/github/stars/josh-silvas/nbot.svg?style=flat-square&label=GitHub%20Stars)](https://github.com/josh-silvas/nbot/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/josh-silvas/nbot.svg?style=flat-square&label=GitHub%20Forks)](https://github.com/josh-silvas/nbot/network/members)
[![GitHub watchers](https://img.shields.io/github/watchers/josh-silvas/nbot.svg?style=flat-square&label=GitHub%20Watchers)](https://github.com/josh-silvas/nbot/watchers)
[![GitHub contributors](https://img.shields.io/github/contributors/josh-silvas/nbot.svg?style=flat-square)](https://github.com/josh-silvas/nbot/graphs/contributors)


***For full documentation, see the [NBot Documentation](https://pages.github.com/josh-silvas/nbot) site***

## Summary
NBot is Nautobot specific CLI tool that allows you to interact with Nautobot via the command line. 
In addition to Nautobot interactions, NBot also provides a set of tools to help with day-to-day, including an underlying ssh
library that allows you to interact with devices via the command line.

Contributions to NBot are extremely welcomed. To contribute, check out our  
[CONTRIBUTING.md](docs/contributing.md) section. 

If you aren't a contributor, but would like to see something new in NBot, 
you can submit a [Github Issue](https://github.com/josh-silvas/nbot/issues) and I or someone from the community would be 
happy to dig into it.

## Usage Guide
NBot uses a shared framework model where the core codebase focuses on system functionality
and key storage, while the `plugin` system offers a wide variety of sub-commands that
utilize the core framework.

![NBot Example #1](docs/images/example_1.gif)

## Installing
NBot is currently built for Mac OSX and Linux 64-bit architectures. If your
OS is not supported and you would like to see it added, please let us know.

See the [INSTALLATION.md](docs/installation.md) guide for installation instructions.

## Contributing
The design behind NBot was intented for it to be easy to contribute to, however, there's
always room for improvement. If you have any suggestions, please feel free to open an issue or 
submit a pull request.

As noted from above, check out the [CONTRIBUTING.md](docs/contributing.md) guide for details on how to contribute. That being said,
I don't want guidelines to get in the way of you contributing, so if you have any questions, please feel free to reach out to me
or submit your pull request and I'll be happy to help you out.


## More help
> Email [Josh Silvas](mailto:josh@jsilvas.com)
