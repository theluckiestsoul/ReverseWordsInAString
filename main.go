// S = i.like.this.program.very.much
// Output: much.very.program.this.like.i
// Explanation: After reversing the whole
// string(not individual words), the input
// string becomes
// much.very.program.this.like.i
package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "i.like.this.program.very.much"
	array := strings.Split(input, ".")

	beg, end := 0, len(array)-1

	for beg < end {
		array[beg], array[end] = array[end], array[beg]
		beg++
		end--
	}

	fmt.Println(strings.Join(array, "."))
}
