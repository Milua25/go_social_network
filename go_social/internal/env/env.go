package env

import (
	"os"
	"strconv"
	"strings"
)

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	// Trim wrapping quotes that can sneak in from .env files or shell exports.
	return strings.Trim(val, `"'`)
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	keyInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return keyInt
}
