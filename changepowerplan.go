package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	var powerplan string
	fmt.Print("Balanced/HighPerformance/Core: ")
	fmt.Scan(&powerplan)
	options := strings.ToLower(powerplan)
	switch options {
	case "balanced":
		balanced()
		fmt.Print("Powerplan balanced")
		break
	case "highperformance":
		highperformance()
		fmt.Print("Powerplan highperformance")
		break
	case "core":
		output, _ := obtener_guid()
		outputcmd := string(output)
		core := strings.Split(outputcmd, "GUID: ")[1]
		scorepowerplan(core)
	}
}

func balanced() ([]byte, error) {
	cmd := exec.Command("cmd", "/C", "powercfg -s 381b4222-f694-41f0-9685-ff5bb260df2e")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return output, nil
}

func highperformance() ([]byte, error) {
	cmd := exec.Command("cmd", "/C", "powercfg -duplicateschema 8c5e7fda-e8bf-4a96-9a85-a6e23a8c635c & powercfg -s 8c5e7fda-e8bf-4a96-9a85-a6e23a8c635c")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return output, nil
}

func obtener_guid() ([]byte, error) {
	cmd := exec.Command("cmd", "/C", "powercfg -import C:\\Core.pow")
	output, _ := cmd.Output()
	return output, nil
}

func scorepowerplan(core string) ([]byte, error) {
	cmd := exec.Command("cmd", "/C", fmt.Sprintf("powercfg -s %s", core))
	output, _ := cmd.Output()
	return output, nil
}
