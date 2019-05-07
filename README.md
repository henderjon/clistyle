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
