package array

import "fmt"

func basic() {
	m := map[string]int{
		"key":  1,
		"test": 2,
	}

	fmt.Println(m)
	fmt.Println(m["key"])

	delete(m, "test")
	fmt.Println(m)
}
