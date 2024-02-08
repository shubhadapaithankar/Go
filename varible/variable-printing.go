//This format string can contain one or more format-specifiers. Some common format specifiers are:
//%d specifies format for integral values.
//%s specifies format for string values.
//%v specifies the general default format.

package main
import "fmt"

var number int = 5  // number declared outside (global scope)

func main(){
    var decision bool =  true // decision declared inside function (local scope)
    fmt.Printf("Original Value of number: %d\n",number)
    number = 10               // reassigning the number
    fmt.Printf("New Value of number: %d\n",number)
    fmt.Printf("Value of decision: %t\n"  ,decision)
  }