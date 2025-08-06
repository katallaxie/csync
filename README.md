# :bank: `csync`

[![Test & Build](https://github.com/katallaxie/csync/actions/workflows/main.yml/badge.svg)](https://github.com/katallaxie/csync/actions/workflows/main.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/katallaxie/csync)](https://goreportcard.com/report/github.com/katallaxie/csync)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)
[![Volkswagen](https://auchenberg.github.io/volkswagen/volkswargen_ci.svg?v=1)](https://github.com/auchenberg/volkswagen)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

`csync` is a tool to backup your settings using common cloud storage solutions like [iCloud Drive](https://www.icloud.com), [Dropbox](https://dropbox.com), [Google Drive](https://www.google.com/intl/de/drive/) or any other file system based sync solution.

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

## Supported Applications

- [x] `alacritty`
- [x] `aws`
- [x] `azure`
- [x] `bash`
- [x] `bartender`
- [x] `docker`
- [x] `bat`
- [x] `ghostty`
- [x] `git`
- [x] `gnupg`
- [x] `brew`
- [x] `hyper`
- [x] `kubectl`
- [x] `macos`
- [x] `magnet`
- [x] `mail`
- [x] `nano`
- [x] `ngrok`
- [x] `npm`
- [x] `raycast`
- [x] `ssh`
- [x] `terminal`
- [x] `tmux`
- [x] `vscode`
- [x] `wget`
- [x] `zed`
- [x] `zsh` 

## Documentation 

[Wiki](https://github.com/katallaxie/csync/wiki)

## Plugins

`csync` supports plugins for `backup`, `restore`, `link` and `unlink` commands.

```go
import "github.com/katallaxie/csync/pkg/plugins/v1"
```

:warning: The support is still under development and the APIs may change in the future.

## Development

The development is intended to be run with [Codespaces](https://github.com/features/codespaces) the blazing fast cloud developer environment.

```bash
env GO111MODULE=on goreleaser release --snapshot --rm-dist
```

## License

[Apache 2.0](/LICENSE)
