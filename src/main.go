package main

import (
	"fmt"
	"log"
	"os/exec"
	"systemsetupgolang/src/packages"
)

func main() {
	fmt.Println("Welcome to System Setup!")
	var command string
	var invalid string = "Invalid Command"
	fmt.Println("Please enter a command.")
	fmt.Println("Your options are, update, install, speedtest.")
	fmt.Scan(&command)
	if command == "update" {
		updateCMD()
	} else if command == "Full Install" {
		updateCMD()
		installCMD()
		packages.SpeedTestCMD()
	} else if command == "install" {
		installCMD()
	} else if command == "speedtest" {
		packages.SpeedTestCMD()
	} else if command == "nginx" {
		packages.InstallNginx()
	} else {
		fmt.Println(invalid)
		fmt.Scanln()
	}
}

func updateCMD() {
	cmd := exec.Command("apt", "update", "-y")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print("There was an error: ", err)
		return
	}
	fmt.Println("Command Output:", string(out))

	upgrade := exec.Command("apt", "upgrade", "-y")
	output, err := upgrade.CombinedOutput()
	if err != nil {
		log.Fatal("Couldnt upgrade", err)
		return
	}
	fmt.Println("Command output:", string(output))
}

func installCMD() {
	packages := []string{"nload", "wireshark", "tshark", "tcpdump", "ipset", "nmap", "net-tools"}
	for _, pkg := range packages {
		cmd := exec.Command("apt", "install", pkg, "-y")
		cmd.Stdin = nil
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("There was an error: ", err)
			return
		}
		fmt.Println("Command Output:", string(out))
	}
}
