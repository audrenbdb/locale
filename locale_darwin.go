//+build darwin

package locale

func getOSLocale() string {
	cmd := exec.Command("sh", "osascript -e 'user locale of (get system info)'")
	output, err := cmd.Output()
	if err != nil {
		return defaultLang
	}
	return output
}