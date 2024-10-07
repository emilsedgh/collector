package util

import "os"

func IsHeroku() bool {
	return false
	return os.Getenv("DYNO") != "" && os.Getenv("PORT") != ""
}
