package main

import (
	"fmt"
	"logoinig/internal/router"
)

func main() {
	fmt.Println("app is running on ':8080' port....")
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("we had error: ", err)
	}
	fmt.Println("exiting...")
}
