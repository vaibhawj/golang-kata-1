package main

import (
	"fmt"

	"github.com/vaibhawj/golang-kata-1/v1/service"
)

func main() {
	fmt.Println(welcomeMessage())

	_, works := service.Load()

	fmt.Println(works)

}

func welcomeMessage() string {
	return "Hello world!"
}
