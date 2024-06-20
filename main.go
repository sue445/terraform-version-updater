package main

import (
	"flag"
	"fmt"
)

var (
	// Version represents app version (injected from ldflags)
	Version string

	// Revision represents app revision (injected from ldflags)
	Revision string

	isPrintVersion bool
)

func printVersion() {
	fmt.Println(GetVersion())
}

func GetVersion() string {
	return fmt.Sprintf("terraform-version-updater %s (revision %s)", Version, Revision)
}

func main() {
	flag.BoolVar(&isPrintVersion, "version", false, "Whether showing version")

	flag.Parse()

	if isPrintVersion {
		printVersion()
		return
	}
}
