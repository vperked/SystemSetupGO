package main

import (
	"fmt"
	"systemsetupgolang/src/packages"
)

func main() {
	fmt.Println("Welcome to System Setup!")
	var command string
	var invalid string = "Invalid Command"
	fmt.Println("Please enter a command.")
	fmt.Println("Your commands are, setup, openvpn, hostname.")
	fmt.Scan(&command)
	switch command {
	case "setup":
		packages.UpdateCMD()
		packages.InstallCMD()
		packages.SpeedTestCMD()
		packages.SetHostname()
		packages.WhiteList("")
	case "openvpn":
		packages.OpenVPN()
	case "hostname":
		packages.SetHostname()
	case "":
		fmt.Println(invalid)
	}
}
