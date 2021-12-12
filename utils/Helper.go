package utils

import "os"

func GetEnvWithKeys(key string) string {
	return os.Getenv(key)
}
