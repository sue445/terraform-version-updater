package updater

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
)

// Updater updates .terraform-version
type Updater struct {
	IsDryRun       bool
	UpdaterVersion string
}

// NewUpdater returns new Updater's instance
func NewUpdater(isDryRun bool, version string) *Updater {
	return &Updater{IsDryRun: isDryRun, UpdaterVersion: version}
}

// Execute performs major processing for updater
func (u *Updater) Execute(targetVersion string, terraformVersionPath string, cooldownDays int) error {
	terraformVersionFile, err := readFile(terraformVersionPath)
	if err != nil {
		return errors.WithStack(err)
	}

	updatedVersionFile, err := UpdateTerraformVersion(&UpdateTerraformVersionParams{
		Src:            terraformVersionFile,
		TargetVersion:  targetVersion,
		UpdaterVersion: u.UpdaterVersion,
		CooldownDays:   cooldownDays,
		CurrentTime:    new(time.Now()),
	})
	if err != nil {
		return errors.WithStack(err)
	}

	if updatedVersionFile == terraformVersionFile {
		u.info(fmt.Sprintf("%s wasn't updated", terraformVersionPath))
		return nil
	}

	if !u.IsDryRun {
		err = os.WriteFile(terraformVersionPath, []byte(updatedVersionFile), 0644)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	beforeVersion := strings.TrimSpace(terraformVersionFile)
	afterVersion := strings.TrimSpace(updatedVersionFile)

	u.info(fmt.Sprintf("%s updated (%s -> %s)", terraformVersionPath, beforeVersion, afterVersion))

	return nil
}

func (u *Updater) info(message string) {
	if u.IsDryRun {
		message += " (dry-run)"
	}
	log.Println(message)
}

func readFile(file string) (string, error) {
	data, err := os.ReadFile(file)

	if err != nil {
		return "", errors.WithStack(err)
	}

	return string(data), nil
}
