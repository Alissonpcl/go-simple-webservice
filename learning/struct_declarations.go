package learning

import "fmt"

type user struct {
	ID        int
	FirstName string
	LastName  string
}

func CreateStructs() {

	var u user
	u.ID = 1
	u.FirstName = "Alisson"
	u.LastName = "Lima"
	fmt.Println(u)

	u2 := user{
		ID: 1,
		FirstName: "Alisson",
		LastName:  "Lima",
	}
	fmt.Println(u2)
}
