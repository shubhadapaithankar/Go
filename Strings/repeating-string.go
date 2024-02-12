// strings.Repeat(s, count int) string

package main

import (
	"fmt"
	"strings"
)

func main() {
	var origS string = "Hi There!"
	var newS string
	newS = strings.Repeat(origS, 3) // Repeating origS 3 times
	fmt.Printf("The new repeated string is: %s\n", newS)
}
