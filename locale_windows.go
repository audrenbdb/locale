//+build windows

package locale

import (
	"os/exec"
	"syscall"
)

func getOSLocale() string {
	cmd := exec.Command("powershell", "Get-Culture | select -exp Name")
	hidePowershellCmd(cmd)
	output, err := cmd.Output()
	if err != nil {
		return defaultLang
	}
	return string(output)
}

func hidePowershellCmd(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000}
}