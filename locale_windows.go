//+build windows

package locale

import (
	"os/exec"
	"strings"
	"syscall"
)

func GetLang() string {
	cmd := exec.Command("powershell", "Get-Culture | select -exp Name")
	hidePowershellCmd(cmd)
	output, err := cmd.Output()
	if err != nil {
		return defaultLang
	}
	return strings.ToLower(strings.Trim(string(output), "\r\n"))
}

func hidePowershellCmd(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000}
}