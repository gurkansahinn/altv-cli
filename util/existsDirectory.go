package util

import "os"

func ExistsDirectory(name string) bool {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}

	return true
}
