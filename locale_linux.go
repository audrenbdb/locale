//+build linux

package locale

import (
	"os"
	"strings"
)

func GetLang() string {
	lang, ok := os.LookupEnv("LANG")
	if !ok {
		return defaultLang
	}
	return strings.ToLower(trimLang(lang))
}

func trimLang(lang string) string {
	return strings.Split(lang, ".")[0]
}
