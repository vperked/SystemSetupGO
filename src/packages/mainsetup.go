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

func WhiteList() {
	fmt.Println("What port would you like to whitelist?")
	var port string
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
}
