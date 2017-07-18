package b

import (
	"sku/test/a"
	"fmt"
)

func init() {
	fmt.Printf("b: %d\n", a.NumA)
}