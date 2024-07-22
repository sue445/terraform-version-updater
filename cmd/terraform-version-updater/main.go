package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/sue445/terraform-version-updater"
	"log"
)

var (
	// Version represents app version (injected from ldflags)
	Version string

	// Revision represents app revision (injected from ldflags)
	Revision string
)

func printVersion() {
	fmt.Println(getVersion())
}

func getVersion() string {
	return fmt.Sprintf("terraform-version-updater %s (revision %s)", Version, Revision)
}

func main() {
	isPrintVersion := pflag.BoolP("version", "v", false, "Whether showing version")
	targetVersion := pflag.StringP("target", "t", "latest", "Version to be updated")
	terraformVersionPath := pflag.StringP("file", "f", ".terraform-version", "Path to .terraform-version file")
	isDryRun := pflag.BoolP("dry-run", "d", false, "Whether dry-run")
	isShowHelp := pflag.BoolP("help", "h", false, "Whether show help")

	pflag.Parse()

	if *isPrintVersion {
		printVersion()
		return
	}

	if *isShowHelp {
		fmt.Println("Usage of terraform-version-updater:")
		pflag.PrintDefaults()
		return
	}

	u := updater.NewUpdater(*isDryRun)
	err := u.Execute(*targetVersion, *terraformVersionPath)
	if err != nil {
		log.Fatal(err)
	}
}
