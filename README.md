# static_analyzer

## できる（やりたい）こと
変数を初期化しないまま使っている箇所を検出します

## 例
```go
package main

import (
  "fmt"
)

func main() {
  var a int
  b := a // want "variable is not initialized" 
  fmt.Println(b)
}
```
