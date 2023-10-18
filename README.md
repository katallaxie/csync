# :bank: `csync`

[![Test & Build](https://github.com/katallaxie/csync/actions/workflows/main.yml/badge.svg)](https://github.com/katallaxie/csync/actions/workflows/main.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/katallaxie/csync)](https://goreportcard.com/report/github.com/katallaxie/csync)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)
[![Volkswagen](https://auchenberg.github.io/volkswagen/volkswargen_ci.svg?v=1)](https://github.com/auchenberg/volkswagen)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

`csync` is a tool to backup your settings using common cloud storage solutions like [iCloud Drive](https://www.icloud.com), [Dropbox](https://dropbox.com), [Google Drive](https://www.google.com/intl/de/drive/) or any other file system based sync solution.

:point_right: [Documentation](https://katallaxie.github.io/csync/)

:warning: This project is under active development and things may change quickly.

## Install

### Homebrew

```bash
brew install katallaxie/csync-tap/csync
```

## Example

```yaml
version: 1
provider:
  name: icloud
```

## Plugins

`csync` supports plugins for `backup`, `restore`, `link` and `unlink` commands. These plugins use [go-plugin](https://github.com/hashicorp/go-plugin) to plug in new features.

:warning: The support is still under development and the APIs may change in the future.

## Development

The development is intended to be run with [Codespaces](https://github.com/features/codespaces) the blazing fast cloud developer environment.

```bash
env GO111MODULE=on goreleaser release --snapshot --rm-dist
```

## License

[Apache 2.0](/LICENSE)
