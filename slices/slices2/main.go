// slices2
// Make me compile!

package main

import "fmt"

func main() {
	names := [4]string{"John", "Maria", "Carl", "Peter"}
    // lastTwoNames := names[] // after figuring out the answer, try with other low/high bounds
    // lastTwoNames := names[2:]  // from idx 2 till the end
    // lastTwoNames := names[2:4]  // from idx 2 till the idx 4-1 [2,4)
    // firstTwoNames := names[:3]  // from idx 0 till the idx 3-1
    lastTwoNames := names[2:]
	fmt.Println(lastTwoNames)
}
