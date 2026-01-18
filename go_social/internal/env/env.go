package env

import (
	"os"
	"strconv"
	"strings"
)

// GetString reads an env var and returns the fallback when missing.
func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	// Trim wrapping quotes that can sneak in from .env files or shell exports.
	return strings.Trim(val, `"'`)
}

// GetInt reads an integer env var or returns the fallback on missing/parse failure.
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

func GetBool(key string, falllback bool) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		return falllback
	}

	boolValue, err := strconv.ParseBool(val)
	if err != nil {
		return falllback
	}
	return boolValue
}
