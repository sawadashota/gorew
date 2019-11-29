Gorew
===
 
Go CLI management tool.
Execute simply `go get` command and record lockfile.


Usage
---

### Add Package

```
$ gorew add github.com/some/package
```

### Update Package

```
$ gorew update github.com/some/package
```

If you want to update all package, give no arguments.

```
$ gorew update
```

### Remove Package

```
$ gorew rm github.com/some/package
```

### List Packages

```
$ gorew list
```


### Install Packages from lockfile

Install packages written in `.gorew`

```
$ gorew install
```

Environment Variables
---

### `GOREW_LOCKFILE_PATH`

Path to `.gorew` file which is lockfile. Default is `~/.gorew`.

Installation
---

```
$ brew tap sawadashota/homebrew-cheers
$ brew install gorew
```

or

```
$ go get -u github.com/sawadashota/gorew
```

License
---

MIT

Author
---

Shota Sawada