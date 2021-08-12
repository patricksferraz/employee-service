package utils

import (
	"bytes"
	"os"
	"unicode"
)

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func CleanNonDigits(str *string) {

	buf := bytes.NewBufferString("")
	for _, r := range *str {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	*str = buf.String()
}
