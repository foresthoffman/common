# ⚙️ Common ⚙️

![Current project build status](https://github.com/foresthoffman/common/actions/workflows/go.yml/badge.svg)

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