package monitor

import (
	"fmt"
	"os/exec"
)

func GetProcessorUsePercentage() (string, error) {
	cmd := exec.Command("powershell", "-Command", "(Get-CimInstance Win32_Processor).LoadPercentage")
	bytes, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(bytes), nil
}
