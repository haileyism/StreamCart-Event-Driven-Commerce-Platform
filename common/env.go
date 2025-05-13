package common

import "syscall"

func EnvString(key, fallback string) string { //receive key and fallback
	if val, ok := syscall.Getenv(key); ok{
		return val

	} //receive the key, return a value and bolean in case not found
	return fallback
}