# Installing
## Mac OSX
To install/update dmux on an OSX build, you can do so using Homebrew. If you
do not already have `brew` installed on your machine, you can find installation
instructions on the [Homebrew website](https://brew.sh/).

* If you have not already, add the dmux tap to your brew taps:
```
brew tap josh-silvas/dmux git@github.com:josh-silvas/dmux
```

* Once tapped, you can install/upgrade/remove dmux using the regular brew methods
```
brew update && brew install dmux

brew update && brew upgrade dmux

brew uninstall dmux
```

## Linux (Debian/Ubuntu)
Installing on Linux is really just pulling down the `.deb` build and running [dpkg](https://man7.org/linux/man-pages/man1/dpkg.1.html)
to install it.

 > TODO: Add instructions for installing on Linux
