package main

import "fmt"

type Beisen struct {
	Name      string
	Age       int
	IsPresent bool
}

type Int int

func structs_example() {
	type User struct {
		name     string
		age      int
		position string
		isAdmin  bool
	}
	user := User{
		name:     "Yerke",
		age:      22,
		position: "mentor",
		isAdmin:  false,
	}
	user2 := User{}
	user2.name = "Beisen"
	user.name = "test"
	fmt.Println(user)
	fmt.Println(user2)
	fmt.Println(user.position)
}
