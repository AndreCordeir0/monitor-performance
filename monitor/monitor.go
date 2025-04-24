package monitor

import (
	"fmt"
	"os/exec"
)

type Monitor struct {
	RamUsage string
	CpuUsage string
}

func NewMonitor() *Monitor {
	return &Monitor{
		RamUsage: "",
		CpuUsage: "",
	}
}

func GetProcessorUsePercentage() (string, error) {
	cmd := exec.Command("powershell", "-Command", "(Get-CimInstance Win32_Processor).LoadPercentage")
	bytes, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(bytes), nil
}

func GetRamUsePercentage() (string, error) {
	cmd := exec.Command("powershell", "-Command", "Get-CimInstance Win32_OperatingSystem | ForEach-Object { [math]::Round((($_.TotalVisibleMemorySize - $_.FreePhysicalMemory) / $_.TotalVisibleMemorySize) * 100) }")
	bytes, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(bytes), nil
}
func (m *Monitor) percent() {
	// m.ramUsage = p
}
