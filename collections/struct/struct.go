package array

import "fmt"

func basic() {
	type User struct {
		id        int
		firstName string
		lastName  string
	}

	var user1 User
	user1.id = 1
	user1.firstName = "Ross"
	user1.lastName = "Moug"

	fmt.Println(user1)
	fmt.Println(user1.id)

	user2 := User{
		id:        1,
		firstName: "Ross",
		lastName:  "Moug",
	}
	fmt.Println(user2)
}
