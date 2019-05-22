package main

import (
	"fmt"
)

func main() {
		a := incrementor()
		b := incrementor()
		fmt.Printf("ANONYMOUS func() int executed. Func location %p, x=%v\n", a, a())
		fmt.Printf("ANONYMOUS func() int executed. Func location %p, x=%v\n", a, a())
		fmt.Printf("ANONYMOUS func() int executed. Func location %p, x=%v\n", a, a())
		fmt.Printf("ANONYMOUS func() int executed. Func location %p, x=%v\n", a, a())
		fmt.Printf("ANONYMOUS func() int executed. Func location %p, x=%v\n", b, b())
		fmt.Printf("ANONYMOUS func() int executed. Func location %p, x=%v\n", b, b())
}

func incrementor() func() int {
		var x int
		fmt.Printf("Inside func incrementor, x=%v\nvar x created at location %p\n", x, &x)
		return func() int {
 			fmt.Printf("Inside ANONYMOUS func() int, enclosed x=%v at location %p\n", x, &x)
			x++
			return x
		}
}