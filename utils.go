package main

import (
	"os"
)

func exists(fn string) bool {
	if _, e := os.Stat(fn); e == nil {
		return true
	}
	return false
}
