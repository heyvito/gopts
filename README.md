# Interspatial Manipulator 👽
<p align="center">
    <a href="https://godoc.org/github.com/victorgama/interman"><img src="https://godoc.org/github.com/victorgama/interman?status.svg" alt="GoDoc"></a>
    <a href="https://travis-ci.org/victorgama/interman"><img src="https://travis-ci.org/victorgama/interman.svg?branch=master" /></a>
    <a href="https://codecov.io/gh/victorgama/interman"><img src="https://codecov.io/gh/victorgama/interman/branch/master/graph/badge.svg" alt="Codecov" /></a>
    <a href="https://goreportcard.com/report/github.com/victorgama/interman"><img src="https://goreportcard.com/badge/github.com/victorgama/interman" /></a>
    <img alt="License" src="https://img.shields.io/badge/license-MIT-blue.svg?style=flat" />
</p>

The Interspatial Manipulator (or `interman`) is a dummy library that loads environment
variables into a given struct.

## Installing

1. Download and install it
```
$ go get -u github.com/victorgama/interman
```

2. Import it in your code:
```go
import "github.com/victorgama/interman"
```

## Usage

This is a really short crash-course. More information can be found on [godoc](https://godoc.org/github.com/victorgama/interman). Assuming an environment with two variables `USERNAME` and `AUTO_RESTART` set to `John` and `false`, respectively, running the following snippet of code:

```go
package main

import (
    "github.com/victorgama/interman"
    "fmt"
)

type Settings struct {
    Username    string
    AutoRestart bool    `default:"true"`
}

func main() {
    settings := interman.LoadEnvs(Settings{}).(Settings)
    fmt.Println(settings.Username)
    fmt.Println(settings.AutoRestart)
}
```

Will yield the following output:

```
John
false
```

## License

```
MIT License

Copyright (c) 2017 Victor Gama

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
