package config

import "os"

func GetEnv(envName string) string {
	if value, exists := os.LookupEnv(envName); exists {
		return value
	}

	return ""
}
