// const identifier [type] = value
const PI = 3.14
const B = "hello"
const B string = "hello"

//Typed and untyped constants#
var n int  
f(n + 5)   // untyped numeric constant "5" becomes typed as int, because n was int.


//Compilation
const C1 = 2/3 //okay
const C2 = getNumber() //not okay : getNumber() can’t provide the value at compile-time

//Overflow
//Numeric constants have no size or sign. They can be of arbitrarily high precision and do not overflow:

const Ln2= 0.693147180559945309417232121458\
176568075500134360255254120680009
const Log2E= 1/Ln2 // this is a precise reciprocal
const BILLION = 1e9 // float constant
const HARD_EIGHT = (1 << 100) >> 97


//Multiple assignments#
const BEEF, TWO, C = "meat", 2, "veg"

const MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY int= 1, 2, 3, 4, 5, 6

//Enumerations#
//Listing of all elements of a set is called enumeration. Constants can be used for enumerations. For example:

const (
	UNKNOWN = 0
	FEMALE = 1
	MALE = 2
  )


  const (
	UNKNOWN = iota
	FEMALE = iota
	MALE = iota
  )


type Gender int
const (
  UNKNOWN Gender = iota
  FEMALE
  MALE
)

//You can give enumeration a type name. For example, FEMALE, MALE and UNKNOWN are categories of Gender. Let’s give them Gender as the type name