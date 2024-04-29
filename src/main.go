package main

import (
	"fmt"
	"log"
	"os/exec"
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
	} else if command == "install" {
		installCMD()
	} else if command == "speedtest" {
		speedTestCMD()
	} else {
		fmt.Println(invalid)
		fmt.Scanln()
	}
}

func updateCMD() {
	cmd := exec.Command("apt", "update", "-y")
	out, err := cmd.Output()
	if err != nil {
		fmt.Print("There was an error: ", err)
	}
	fmt.Println("Command Output:", string(out))
	upgrade := exec.Command("apt", "upgrade", "-y")
	output, err := upgrade.Output()
	if err != nil {
		log.Fatal("Couldnt upgrade", err)
	}
	fmt.Println("Command output:", string(output))
}

func installCMD() {
	packages := []string{"nload", "wireshark", "tshark", "tcpdump", "ipset", "nmap", "net-tools"}
	for _, pkg := range packages {
		cmd := exec.Command("apt", "install", pkg, "-y")
		cmd.Stdin = nil
		out, err := cmd.Output()
		if err != nil {
			fmt.Println("There was an error: ", err)
			continue
		}
		fmt.Println("Command Output:", string(out))
	}
}

func speedTestCMD() {
	cmd := exec.Command("bash", "-c", "curl -s https://packagecloud.io/install/repositories/ookla/speedtest-cli/script.deb.sh | sudo bash && sudo apt-get install speedtest -y")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("There was an error: ", err)
	}
	fmt.Println("Command Output: ", string(out))
}