package main

import "fmt"

func primitives() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	s := "string"
	fmt.Println(s)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)
}

func pointers() {
	var s *string = new(string)

	*s = "string"

	fmt.Println(s)
	fmt.Println(*s)

	var x string = "text"
	fmt.Println(x)

	p := &x
	fmt.Println(p, *p)

	x = "abc"
	fmt.Println(p, *p)
}
