package updater

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Updater updates .terraform-version
type Updater struct {
	IsDryRun bool
}

// NewUpdater returns new Updater's instance
func NewUpdater(isDryRun bool) *Updater {
	return &Updater{IsDryRun: isDryRun}
}

// Execute performs major processing for updater
func (u *Updater) Execute(targetVersion string, terraformVersionPath string) error {
	terraformVersionFile, err := readFile(terraformVersionPath)
	if err != nil {
		return err
	}

	versions, err := GetTerraformStableVersions()
	if err != nil {
		return err
	}

	updatedVersionFile := UpdateTerraformVersion(&UpdateTerraformVersionParams{
		Src:           terraformVersionFile,
		TargetVersion: targetVersion,
		Versions:      versions,
	})

	if updatedVersionFile == terraformVersionFile {
		u.Info(fmt.Sprintf("%s wasn't updated", terraformVersionPath))
		return nil
	}

	if !u.IsDryRun {
		err = os.WriteFile(terraformVersionPath, []byte(updatedVersionFile), 0644)
		if err != nil {
			return err
		}
	}

	beforeVersion := strings.TrimSpace(terraformVersionFile)
	afterVersion := strings.TrimSpace(updatedVersionFile)

	u.Info(fmt.Sprintf("%s updated (%s -> %s)", terraformVersionPath, beforeVersion, afterVersion))

	return nil
}

func (u *Updater) Info(message string) {
	if u.IsDryRun {
		message += " (dry-run)"
	}
	log.Println(message)
}

func readFile(file string) (string, error) {
	data, err := os.ReadFile(file)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
