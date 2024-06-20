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

	targetVersion string

	terraformVersionPath string
)

func printVersion() {
	fmt.Println(GetVersion())
}

func GetVersion() string {
	return fmt.Sprintf("terraform-version-updater %s (revision %s)", Version, Revision)
}

func init() {
	flag.BoolVar(&isPrintVersion, "version", false, "Whether showing version")
	flag.StringVar(&targetVersion, "target", "latest", "Version to be updated")
	flag.StringVar(&terraformVersionPath, "file", ".terraform-version", "Path to .terraform-version file")
}

func main() {
	flag.Parse()

	if isPrintVersion {
		printVersion()
		return
	}
}
