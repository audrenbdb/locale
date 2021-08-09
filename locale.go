package locale

import "strings"

const (
	defaultLang = "en"
)

func GetLang() string {
	l := getOSLocale()
	l = strings.TrimSpace(l)
	l = strings.ToLower(l)
	return l[0:2]
}
