# DMux : Nautobot CLI Tool 
[![Go Report Card](https://goreportcard.com/badge/github.com/josh-silvas/dmux)](https://goreportcard.com/report/github.com/josh-silvas/dmux)
[![GoDoc](https://godoc.org/github.com/josh-silvas/dmux?status.svg)](https://godoc.org/github.com/josh-silvas/dmux)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/josh-silvas/dmux/blob/main/LICENSE)
[![Release](https://img.shields.io/github/release/josh-silvas/dmux.svg?style=flat-square)](https://github.com/josh-silvas/dmux/releases/latest)
[![Downloads](https://img.shields.io/github/downloads/josh-silvas/dmux/total.svg?style=flat-square)](https://github.com/josh-silvas/dmux/releases)
[![GitHub issues](https://img.shields.io/github/issues/josh-silvas/dmux.svg?style=flat-square)](https://github.com/josh-silvas/dmux/issues)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/josh-silvas/dmux.svg?style=flat-square)](https://github.com/josh-silvas/dmux/pulls)
[![GitHub contributors](https://img.shields.io/github/contributors/josh-silvas/dmux.svg?style=flat-square)](https://github.com/josh-silvas/dmux/graphs/contributors)
[![GitHub stars](https://img.shields.io/github/stars/josh-silvas/dmux.svg?style=flat-square&label=GitHub%20Stars)](https://github.com/josh-silvas/dmux/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/josh-silvas/dmux.svg?style=flat-square&label=GitHub%20Forks)](https://github.com/josh-silvas/dmux/network/members)
[![GitHub watchers](https://img.shields.io/github/watchers/josh-silvas/dmux.svg?style=flat-square&label=GitHub%20Watchers)](https://github.com/josh-silvas/dmux/watchers)
[![GitHub contributors](https://img.shields.io/github/contributors/josh-silvas/dmux.svg?style=flat-square)](https://github.com/josh-silvas/dmux/graphs/contributors)


***For full documentation, see the [DMux Documentation](https://pages.github.com/josh-silvas/dmux) site***

## Summary
DMux is Nautobot specific CLI tool that allows you to interact with Nautobot via the command line. 
In addition to Nautobot interactions, DMux also provides a set of tools to help with day-to-day, including an underlying ssh
library that allows you to interact with devices via the command line.

Contributions to DMux are extremely welcomed. To contribute, check out our  
[CONTRIBUTING.md](docs/contributing.md) section. 

If you aren't a contributor, but would like to see something new in DMux, 
you can submit a [Github Issue](https://github.com/josh-silvas/dmux/issues) and I or someone from the community would be 
happy to dig into it.

## Usage Guide
DMux uses a shared framework model where the core codebase focuses on system functionality
and key storage, while the `plugin` system offers a wide variety of sub-commands that
utilize the core framework.

![DMux Example #1](docs/images/example_1.gif)

## Installing
DMux is currently built for Mac OSX and Linux 64-bit architectures. If your
OS is not supported and you would like to see it added, please let us know.

See the [INSTALLATION.md](docs/installation.md) guide for installation instructions.

## Contributing
The design behind DMux was intented for it to be easy to contribute to, however, there's
always room for improvement. If you have any suggestions, please feel free to open an issue or 
submit a pull request.

As noted from above, check out the [CONTRIBUTING.md](docs/contributing.md) guide for details on how to contribute. That being said,
I don't want guidelines to get in the way of you contributing, so if you have any questions, please feel free to reach out to me
or submit your pull request and I'll be happy to help you out.


## More help
> Email [Josh Silvas](mailto:josh@jsilvas.com)
