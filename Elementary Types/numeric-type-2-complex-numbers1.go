package main
import "fmt"

func main(){
    var c1 complex64 = 5 + 10i    // Declaring complex num (real +imaginary(¡))
    fmt.Printf("The value is: %v", c1)
}


//Complex numbers
//A complex number is written in the form of:
//re + im¡

//where re is the real part, im is the imaginary part, and ¡ is the √ -1. For these data we have the following types:
//complex64 (with a 32 bit real and imaginary part each)
//complex128 (with a 64 bit real and imaginary part each)