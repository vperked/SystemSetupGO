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
