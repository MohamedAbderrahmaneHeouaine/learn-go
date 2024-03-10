package main

import (
	"fmt"
	"newproject/types"
)

func main() {
	user := types.User{
		Username: "aymen",
		Age:      33,
	}
	fmt.Println("user")
	fmt.Println(user)
}
