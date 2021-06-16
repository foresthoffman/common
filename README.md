# ⚙️ Common ⚙️

![Current project build status on Ubuntu 18.04](https://github.com/foresthoffman/common/actions/workflows/go-ubuntu-18.04.yml/badge.svg)
![Current project build status on Windows Server 2019](https://github.com/foresthoffman/common/actions/workflows/go-windows-2019.yml/badge.svg)
![Current project build status on Ubuntu Server Latest](https://github.com/foresthoffman/common/actions/workflows/go-windows-latest.yml/badge.svg)
![Current project build status on MacOS 11.0](https://github.com/foresthoffman/common/actions/workflows/go-macos-11.0.yml/badge.svg)
![Current project build status on MacOS Latest](https://github.com/foresthoffman/common/actions/workflows/go-macos-latest.yml/badge.svg)

`common` provides helpful DRY logic common to most Go services.

### Installation

Run `go get -u github.com/foresthoffman/common`

If you're using `go mod`, run `go mod vendor` afterwards.

### Importing

Import this package by including `github.com/foresthoffman/common` in your import block.

e.g.

```go
package main

import(
    ...
    "github.com/foresthoffman/common"
)
```

### Testing

Run `go test -v -count=1 ./...` in the project root directory. Use the `-count=1` to force the tests to run un-cached.

_That's all, enjoy!_