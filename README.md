gorew
===

Golang cli management tool.

Usage
---

### Initialization

Save golang cli to `~/.gorew` by compared `$GOPATH/bin/` and under `$GOPATH/src`.

```bash
gorew init [-b PATH_TO_BIN_IN_GOLANG] [-s PATH_TO_GO_SRC] [-f PATH_TO_SAVE_.gorew]
```

### Install Cli

```bash
gorew install [-f PATH_TO_.gorew]
```

Installation
---

```bash
brew tap sawadashota/homebrew-cheers
brew install gorew
```

or

```bash
go get -u github.com/sawadashota/gorew
```

License
---

MIT

Author
---

Shota Sawada