package ports

import (
	"os/exec"
	"strings"
	"syscall"
)

func ExcludedRanges() ([]string, error) {
	cmd := exec.Command("netsh", "int", "ipv4", "show", "excludedportrange", "protocol=tcp")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var ranges []string
	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimSpace(line)
		parts := strings.Fields(line)
		if len(parts) >= 2 && isNumeric(parts[0]) && isNumeric(parts[1]) {
			ranges = append(ranges, parts[0]+"-"+parts[1])
		}
	}
	return ranges, nil
}

func isNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}