---
title: Overview
description: csync is a tool to backup your settings using common cloud storage solutions like iCloud Drive, Dropbox, Google Drive or any other file system based sync solution.
---

`csync` is a tool to backup your settings using common cloud storage solutions like [iCloud Drive](https://www.icloud.com), [Dropbox](https://dropbox.com), [Google Drive](https://www.google.com/intl/de/drive/) or any other file system based sync solution.

[Installation](/docs/installation) is super easy. You can initialize a config via `csync init` your home directory.

```yaml
# .csync.yml
version: 1
provider:
  name: icloud
```

See [Reference](reference) to learn about additional features and the specficiation for the `.csync.yml` configuration file.