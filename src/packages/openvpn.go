package packages

import (
	"fmt"
	"os/exec"

	"github.com/rs/zerolog/log"
)

func OpenVPN() {
	cmd := exec.Command("curl", "-O", "https://raw.githubusercontent.com/angristan/openvpn-install/master/openvpn-install.sh")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Err(err)
	}
	fmt.Println(string(out))
	chiMod := exec.Command("chmod", "+x", "openvpn-install.sh")
	cmdOut, err := chiMod.CombinedOutput()
	if err != nil {
		log.Err(err)
	}
	fmt.Println(string(cmdOut))

	finish := exec.Command("./openvpn-install.sh")
	finishOut, err := finish.CombinedOutput()
	if err != nil {
		log.Err(err)
	}
	fmt.Println(string(finishOut))
}
