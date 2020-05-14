# opencc

Simplified version of [gocc](https://github.com/liuzl/gocc).

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
