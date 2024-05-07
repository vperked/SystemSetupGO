package packages

import (
	"fmt"
	"log"
	"os/exec"
)

func SetHostname() {
	fmt.Println("What would you like your hostname to be?")
	var hostname string
	fmt.Scan(&hostname)
	cmd := exec.Command("hostnamectl", "set-hostname", hostname)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
	fmt.Println("Set to:", hostname)
}

func WhiteList(port string) {
	fmt.Println("What port would you like to whitelist?")
	var ip string
	fmt.Scan(&port)
	fmt.Println("What IP?")
	fmt.Scan(&ip)
	cmd := exec.Command("iptables", "-A", "INPUT", "-p", "tcp", "-s", ip, "--dport", port, "-j", "ACCEPT")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
	fmt.Println("Would you like to drop the port?")
	var drop string
	fmt.Scan(&drop)
	if drop == "yes" {
		dropPort("")
	} else {
		return
	}
}

func dropPort(port string) {
	fmt.Println("What port would you like to drop?")
	fmt.Scan(&port)
	cmd := exec.Command("iptables", "-A", "INPUT", "-p", "tcp", "--dport", port, "-j", "DROP")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)

}

func UpdateCMD() {
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

func InstallCMD() {
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
