package constant

import "fmt"

const (
	pi    = 3.1415
	first = iota
	second
)

const (
	third = iota
	fourth
)

func basic() {
	fmt.Println(pi)

	const c = 3
	fmt.Println(c + 3)
	fmt.Println(c + 1.2)
}

func iotaConstants() {
	fmt.Println(pi)

	fmt.Println(first, second, third, fourth)
}
