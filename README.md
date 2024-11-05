# opencc

Simplified version of [gocc](https://github.com/liuzl/gocc) and [OpenCC](https://github.com/BYVoid/OpenCC).

OpenCC in one executable. You can also use OpenCC in browser.
See [in-browser Chinese conversion tool](https://caiguanhao.github.io/opencc/wasm/).

The config and dictionary Go files are generated from the source of OpenCC version 1.1.8.
You can also generate them using `go test`.

## Command Line

```
go install github.com/caiguanhao/opencc/opencc@latest
```

## Import Package

### All

```go
package main

import (
	"fmt"

	"github.com/caiguanhao/opencc"
)

func main() {
	// fmt.Println(opencc.Dictionaries)
	fmt.Println(opencc.Convert("s2twp", `鼠标里面的硅二极管坏了，导致光标分辨率降低。`))
	// 滑鼠裡面的矽二極體壞了，導致游標解析度降低。
}
```

### Specific

```go
package main

import (
	"fmt"

	"github.com/caiguanhao/opencc/configs/tw2sp"
)

func main() {
	fmt.Println(tw2sp.Description)
	// Traditional Chinese (Taiwan standard) to Simplified Chinese (with phrases)
	fmt.Println(tw2sp.Dicts.Convert(`滑鼠裡面的矽二極體壞了，導致游標解析度降低。`))
	// 鼠标里面的硅二极管坏了，导致光标分辨率降低。
}
```
