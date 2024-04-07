package main

import (
	"fmt"
	"os/exec"
)

func main() {
	userInput()
}

func userInput() {
	var command string
	fmt.Println("Please enter a command.")
	fmt.Scan(&command)
	if command == "yo" {
		test()
	} else {
		fmt.Println("Invalid command.")
	}
}

func test() {
	cmd := exec.Command("ls")
	err := cmd.Run()
	if err != nil {
		fmt.Println("There was an error, ", err)
	}
}
