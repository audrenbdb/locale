//+build darwin

package locale

import "strings"

func GetLang() string {
	cmd := exec.Command("sh", "osascript -e 'user locale of (get system info)'")
	output, err := cmd.Output()
	if err != nil {
		return defaultLang
	}
	langLocRaw := strings.TrimSpace(string(output))
	return strings.ToLower(strings.Split(langLocRaw, "_")[0])
}