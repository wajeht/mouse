
# 🐁 Mouse (Go Edition)

![go](https://img.shields.io/badge/Go-v1.22-blue)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/wajeht/mouse/blob/main/LICENSE) [![Open Source Love svg1](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/wajeht/mouse)

Move mouse in a square — why? why not?

## 🛠️ Installation

> [!WARNING]
> Before installing **mouse**, make sure Go is installed (version 1.20+ recommended).
> On macOS, grant terminal apps accessibility permissions via
> `System Settings > Privacy & Security > Accessibility`.

```bash
$ go install github.com/wajeht/mouse@latest
```

Make sure your `GOBIN` (default: `~/go/bin`) is in your `PATH`:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"

```

## 🚀 Usage

```bash
$ mouse
# Moving the mouse in a square..., (press Ctrl + C to stop)
$ mouse --verbose
# Moving the mouse in a square..., (press Ctrl + C to stop)
# [verbosemod] Moving right...
# [verbosemod] Moving down...
# ...
```

## © License

Distributed under the MIT License © wajeht. See [LICENSE](./LICENSE) for more information.

