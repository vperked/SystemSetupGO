package packages

import (
	"fmt"
	"os/exec"

	"github.com/rs/zerolog/log"
)

func InstallNginx() {
	cmd := exec.Command("apt", "install", "nginx", "-y")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Err(err)
	}
	fmt.Println(string(out))
}
