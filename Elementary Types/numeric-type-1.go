//Boolean type
//The possible values of this type are the predefined constants true and false. For example:
//var b bool = true

package main
import "fmt"

func main() {
    var n int16 = 34    // int16 variable
    var m int32         // int32 variable

    m = int32(n)        // explicit typing
    fmt.Printf("32 bit int is: %d\n", m)
    fmt.Printf("16 bit int is: %d\n", n)
}