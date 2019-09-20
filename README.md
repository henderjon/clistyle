[![GoDoc](https://godoc.org/github.com/henderjon/clistyle?status.svg)](https://godoc.org/github.com/henderjon/clistyle)
[![License: BSD-3](https://img.shields.io/badge/license-BSD--3-blue.svg)](https://img.shields.io/badge/license-BSD--3-blue.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/henderjon/clistyle)](https://goreportcard.com/report/github.com/henderjon/clistyle)
[![Build Status](https://travis-ci.org/henderjon/clistyle.svg?branch=dev)](https://travis-ci.org/henderjon/clistyle)
![tag](https://img.shields.io/github/tag/henderjon/clistyle.svg)
![release](https://img.shields.io/github/release/henderjon/clistyle.svg)


# clistyle
A very simple way to colorize cli output.

## usage

Named attributes are passed (or'ed) akin to passing flags to the `log` package.

```go
package foo

import (
	"fmt"
	cli "github.com/henderjon/clistyle"
)

func main() {
	s := cli.Style("bar", Bold|YellowBG)
	fmt.Println(s)
}
```
