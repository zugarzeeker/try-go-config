package config

import "os"

func GetEnv(key string, args ...interface{}) interface{} {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	if len(args) != 0 {
		return args[0]
	}
	return nil
}
