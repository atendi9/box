# box

`box` is a minimalist, type-safe utility package for Go that introduces generic container types to safely handle the presence or absence of data without relying on risky `nil` pointer dereferences.

By leveraging Go generics, `box` provides a clean, expressive API inspired by functional programming patterns like `Optional` and `Maybe`.


## 🚀 Quick Start

```go
package main

import (
	"fmt"
	"github.com/atendi9/box"
)

func main() {
	// Creating a present value
	name := box.NewSome("Capivara")
	if name.IsPresent() {
		fmt.Println("Hello,", name.Get()) // Output: Hello, Capivara
	}

	// Creating an empty value
	empty := box.NewNone[string]()
	fmt.Println("Value:", empty.Get()) // Output: Value: "" (returns zero value safely)
}