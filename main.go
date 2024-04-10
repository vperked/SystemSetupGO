package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var command string
	var invalid string = "Invalid Command"
	fmt.Println("Please enter a command.")
	fmt.Scan(&command)
	if command == "root" {
		enableRootCMD()
	} else if command == "update" {
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
		fmt.Print("There was an error: ", err)
	}
	fmt.Println("Command Output:", string(out))
}

func installCMD() {
	packages := []string{"nload", "wireshark", "tshark", "tcpdump", "ipset", "nmap", "net-tools"}
	for _, pkg := range packages {
		cmd := exec.Command("apt", "install", pkg)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("There was an error: ", err)
		}
		fmt.Println("Command Output:", string(out))
	}

}

func speedTestCMD() {
	cmd := exec.Command("bash", "-c", "curl -s https://packagecloud.io/install/repositories/ookla/speedtest-cli/script.deb.sh | sudo bash && sudo apt-get install speedtest")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("There was an error: ", err)
	}
	fmt.Println("Command Output: ", string(out))
}

func enableRootCMD() {
	cmd := exec.Command("sed -i", "'s/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	fmt.Println("Command Output: ", string(out))
}
