// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	L "test/lib"
)

// type myDefined string

func main() {
	// var fname = "Korni"
	//var lname = "Srinivas"
	//var a int = 10000000000000000
	//var b int = 100001
	//var z = 22234567898765456
	//var k = 876543234567897654
	//var x float32 = 20.99
	//c := 3.11111234
	//d := 9.345678
	//f := c * d
	//var y float32 = 19.32
	//fmt.Printf("Hello %s %s \n", lname, fname)
	//fmt.Printf("%d * %d =%d \n", a, b, a*b)
	//fmt.Printf("%g * %g =%g \n", x, y, x*y)
	//fmt.Printf("%T\n", f)
	//fmt.Printf("%T\n", lname)
	//fmt.Printf("%T\n", a)
	//fmt.Printf("%T\n", c)
	//fmt.Printf("%T\n", z*k)

	//---------------------String--------------------

	// var def myDefined
	// var s string = "Srinivas"
	// fmt.Println(len(s))
	// fmt.Printf("%T", def)
	var a int = 20
	var b int = 2
	//------------blank identifiers functions-------------------
	add, _, _, _ := L.Cal(a, b)
	_, mul, _, _ := L.Cal(a, b)
	_, _, div, _ := L.Cal(a, b)
	_, _, _, sub := L.Cal(a, b)

	fmt.Println(add)
	defer fmt.Println(mul)
	defer fmt.Println(div)
	fmt.Println(sub)

}
