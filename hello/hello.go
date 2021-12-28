package main

import (
	"fmt"

	"example.com/greetings"
)

func Tmain() {
	message := greetings.Hello()
	fmt.Println(message)
}
