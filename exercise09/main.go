package main

import "fmt"

/*
create the following variables with the following scopes:
○ package level
■ create outside of func main
■ use the
● var keyword
Todd McLeod - Learn To Code Go - Page 39
● const keyword
○ block level
■ inside func main
■ use the short declaration operator
● use the variables in func main
*/
var x = 40

const y = 41

func main() {
	z := 42
	fmt.Printf("The value of x is %v and the ype of x is %T\n", x, x)
	fmt.Printf("The value of y is %v and the ype of y is %T\n", y, y)
	fmt.Printf("The value of z is %v and the ype of z is %T\n", z, z)

}
