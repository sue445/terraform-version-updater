package main

import (
	"flag"
	"fmt"
	"github.com/sue445/terraform-version-updater"
	"log"
)

var (
	// Version represents app version (injected from ldflags)
	Version string

	// Revision represents app revision (injected from ldflags)
	Revision string

	isPrintVersion bool

	targetVersion string

	terraformVersionPath string

	isDryRun bool
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
	flag.BoolVar(&isDryRun, "dry-run", false, "Whether dry-run")
}

func main() {
	flag.Parse()

	if isPrintVersion {
		printVersion()
		return
	}

	u := updater.NewUpdater(isDryRun)
	err := u.Execute(targetVersion, terraformVersionPath)
	if err != nil {
		log.Fatal(err)
	}
}
