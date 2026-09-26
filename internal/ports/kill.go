package ports

import (
	"fmt"
	"os/exec"
	"syscall"
)

func KillByPort(port uint32, force bool) error {
	list, err := List()
	if err != nil {
		return err
	}
	for _, p := range list {
		if p.Port == port {
			args := []string{"/PID", fmt.Sprint(p.PID)}
			if force {
				args = append(args, "/F")
			}
			cmd := exec.Command("taskkill", args...)
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			if err := cmd.Run(); err != nil {
				return fmt.Errorf("taskkill failed: %w", err)
			}
			return nil
		}
	}
	return fmt.Errorf("no listening process found on port %d", port)
}