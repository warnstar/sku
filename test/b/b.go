package b

import (
	"fmt"
	"sku/test/a"
)

func init() {
	fmt.Printf("b: %d\n", a.NumA)
}
