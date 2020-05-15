# opencc

Simplified version of [gocc](https://github.com/liuzl/gocc) and [OpenCC](https://github.com/BYVoid/OpenCC).

OpenCC in one executable. You can also use OpenCC in browser.
See [in-browser Chinese conversion tool](https://caiguanhao.github.io/opencc/wasm/).

## Command Line

```
go get -v -u github.com/caiguanhao/opencc/opencc
```

## Import Package

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
