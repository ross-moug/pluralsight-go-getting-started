package main

import (
	"fmt"

	"github.com/ross-moug/pluralsight-go-getting-started/models"
)

func main() {
	user := models.User{
		Id:        2,
		FirstName: "Ross",
		LastName:  "Moug",
	}

	fmt.Println(user)
}
