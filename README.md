# SIV (Stable Index Vector) in Go

Implementation of the Stable Index Vector data structure in Go. This structure allows O(1) (constant) time complexity for Add, Remove and Get operations.

## Quick start

### Common usage

```go
package main

import (
    "fmt"
    "github.com/frpelaez/siv"
)

func main() {
    fruits := siv.New(10)

    h1 := fruits.Add("banana")
    h2 := fruits.Add("apple")
    h3 := fruits.Add("kiwi")

    fruits.Remove(h2)

    val, ok := fruits.Get(h2)
    if !ok {
        fmt.Println("Used an invalidated handle to acces an item")
    } else {
        fmt.Print(*val)
    }

    h2 = fruits.Add(orange)

    fmt.Println("-- Conntents --")
    fruits.ForEach(func(h Handle, v *string) bool {
        fmt.Printf("Index: %d (Generation: %d), Value: %s\n", h.Index, h.Generation, *v)
        return true
    })
}
```

### Output

```py
# Used an invalidated handle to acces an item
# -- COntents --
# Index: 0 (Generation: 0), Value: "banana"
# Index: 1 (Generation: 1), Value: "orange"
# Index: 2 (Generation: 0), Value: "apple"
# Index: 3 (Generation: 0), Value: "kiwi"
```
