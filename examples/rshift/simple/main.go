package main

import (
	"fmt"
	"math/bits"

	"github.com/openacid/must"
)

func rshift(a, b int) int {

	// Simple check:
	//   "go build" emits a No-op instruction.
	//   "go build -tags debug" will check "b" not to be zero.
	must.Be.NotZero(b)

	// Inappropriate for complex check:
	// argument expressions are still evaluated.
	must.Be.True(bits.TrailingZeros(uint(a)) > 2)

	return a >> uint(b)
}

func main() {
	// panic at line 20 with "go run -tags debug"
	fmt.Println(rshift(0xf, 1))
}
