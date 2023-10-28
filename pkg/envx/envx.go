package envx

import (
	"os"
	"strconv"
	"strings"
)

func GetString(key string, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultVal
	}

	return value
}

func GetArray(key string, defaultVal string) []string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultVal
	}
	return strings.Split(value, ",")
}

func GetInt(key string, defaultVal int) int {
	valStr := os.Getenv(key)
	value, err := strconv.ParseInt(valStr, 10, 32)
	if err != nil {
		return defaultVal
	}

	return int(value)
}
