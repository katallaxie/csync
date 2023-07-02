---
title: Reference
summary: Reference for `csync` configuration and usage.
---

## CLI

`csync` has the following command line syntax:

```bash
csync [--flags] [-- ARGS...]
```

| Short | Flag | Type | Default | Description |
| - | - | - | - | - |
| `-c` | `--config` | `string` | `.csync.yml` | Config file. Enabled by default. Set to `$HOME/.csync.yml` or change to the location of your config |
| `-r` | `--restore` | `bool` | `false` | Restores files from the cloud storage and creates the symbolic links for syncing |
| `-f` | `--force` | `bool` | `false` | Forces the execution of operations. |
| `-v` | `--verbose` | `bool` | `false` | Enables verbose logging of runtime information. |
| `-d` | `--dry` | `bool` | `false` | Does not apply destructive operations. |
|  | `--unlink` | `bool` | `false` | Restores files from the cloud storage without creating symbolic links |
|  | `--validate` | `bool` | `false` | Validates the specification file provided via `.run.yml`. |
|  | `--init` | `bool` | `false` | Creates a new `.csync.yml` file at the provided location of `--config` (default: `$CWD/.csync.yml`) |
|  | `--version` | `bool` | `false` | Prints the current version. |

## Schema

### Example

```yaml
version: 1
provider:
  name: icloud
```

### General

| Attribute | Type | Default | Description |
| - | - | - | - |
| `version` | `int` | | Specification version to be used. The current version is `1`. |
| `provider` | [`Provider`](#provider) | | Applications to sync. |
| `apps` | [`Apps`](#app) | | Applications to sync. |
| `includes` | `[]string` | | Overwrites any files listes in `apps` to include. |
| `excludes` | `[]string` | | Overwrites any files listes in `apps` to exclude. |

### Provider

| Attribute | Type | Default | Description |
| - | - | - | - |
| `name` | `string` | | Name of the provider `icloud`, `dropbox`, `files`. |
| `path` | `string` | | Path to the provider folder. Only with `files` provider |
| `directory` | `string` | `csync` | The directory in the provider to locate the syncing files. |

### Apps

| Attribute | Type | Default | Description |
| - | - | - | - |
| `name` | `string` | | Name of the application. |
| `files` | `[]stirng` | | List of files to sync. References to the home directory (`~`) are expanded to absolute paths. |