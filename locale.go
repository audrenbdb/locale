package locale

import "strings"

const (
	defaultLang = "en"
)

//GetLang returns first two letters of runtime language
//or "en" if not found / error.
func GetLang() string {
	l := getOSLocale()
	l = strings.TrimSpace(l)
	l = strings.ToLower(l)
	return l[0:2]
}
