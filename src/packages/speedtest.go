package packages

import (
	"fmt"
	"os/exec"
)

func SpeedTestCMD() {
	cmd := exec.Command("bash", "-c", "curl -s https://packagecloud.io/install/repositories/ookla/speedtest-cli/script.deb.sh | sudo bash && sudo apt-get install speedtest -y")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("There was an error: ", err)
	}
	fmt.Println("Command Output: ", string(out))
}
