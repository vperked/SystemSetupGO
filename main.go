package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var command string
	fmt.Println("Please enter a command.")
	fmt.Scan(&command)

	if command == "reboot" {
		rebootCMD()
	} else if command == "update" {
		updateCMD()
	} else {
		fmt.Println("Invalid Command.")
	}
}

func rebootCMD() {
	cmd := exec.Command("reboot")
	err := cmd.Run()
	if err != nil {
		fmt.Println("There was an error, ", err)
	}
}

func updateCMD() {
	cmd := exec.Command("apt", "update")
	out, err := cmd.Output()
	if err != nil {
		fmt.Print("There was an error, ", err)
	}
	fmt.Println("Command Output:", string(out))
}
