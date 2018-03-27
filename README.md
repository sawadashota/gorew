gocmd
===

Golang cli management tool.

Usage
---

### Initialization

Save golang cli to `~/.gocmd` by compared `$GOPATH/bin/` and under `$GOPATH/src`.

```bash
gocmd init [-b PATH_TO_BIN_IN_GOLANG] [-s PATH_TO_GO_SRC] [-f PATH_TO_SAVE_.gocmd]
```

### Install Cli

```bash
gocmd install [-f PATH_TO_.gocmd]
```

Installation
---

```bash
brew tap sawadashota/homebrew-gocmd
brew install gocmd
```

or

```bash
go get -u github.com/sawadashota/gocmd
```

License
---

MIT

Author
---

Shota Sawada