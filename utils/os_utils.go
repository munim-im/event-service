package utils

import "os"
// helper functions for checking environment variables, if not found provide default value
func GetEnvironmentVariable(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
