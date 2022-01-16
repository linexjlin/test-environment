package testenv

import (
	"os"
)

var (
	TEST  = false
	DEGUB = false
)

func init() {
	if debug := os.Getenv("Debug") + os.Getenv("debug") + os.Getenv("DEBUG"); debug != "" {
		DEGUB = true
	}
	if test := os.Getenv("Test") + os.Getenv("test") + os.Getenv("TEST"); test != "" {
		TEST = true
	}
}
