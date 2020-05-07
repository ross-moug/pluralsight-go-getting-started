package array

import "fmt"

func fixed() {
	array := [3]int{1, 2, 3}

	slice := array[:]

	array[0] = 42
	slice[0] = 21

	fmt.Println(array)
	fmt.Println(slice)
}

func nonFixed() {
	slice := []int{1, 2, 3}

	fmt.Println(slice)

	slice = append(slice, 4, 5, 6)

	fmt.Println(slice)

	slice2 := slice[1:]
	slice3 := slice[:2]
	slice4 := slice[1:2]

	fmt.Println(slice2, slice3, slice4)
}
