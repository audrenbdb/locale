//+build linux

package locale

import (
	"os"
)

func getOSLocale() string {
	lang, ok := os.LookupEnv("LANG")
	if !ok {
		return defaultLang
	}
	return lang
}


