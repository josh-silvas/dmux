# DMux : NetEng CLI Tool
* [Summary](#summary)
* [Installation](#installing)
* [Usage Guide](#usage-guide)
* [Contributing](#contributing)
* [More Help](#more-help)

## Summary
DMux is a CLI tool that allows you to interact with Nautobot, NetBox, and potentially any SoT via the command line.
In addition to SoT interactions, DMux also provides a set of tools to help with day-to-day, including an underlying ssh
library that allows you to interact with devices via the command line.

Contributions to DMux are extremely welcomed. To contribute, check out our  
[CONTRIBUTING.md](docs/contributing.md) section.

If you aren't a contributor, but would like to see something new in DMux,
you can submit a [GitHub Issue](https://github.com/josh-silvas/dmux/issues) and I or someone from the community would be
happy to dig into it.

## Usage Guide
DMux uses a shared framework model where the core codebase focuses on system functionality
and key storage, while the `plugin` system offers a wide variety of sub-commands that
utilize the core framework.

![DMux Example #1](images/example_1.gif)

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
