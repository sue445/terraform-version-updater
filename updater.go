package updater

import (
	"log"
	"os"
	"strings"
)

// Execute performs major processing for updater
func Execute(targetVersion string, terraformVersionPath string) error {
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
		log.Printf("%s wasn't updated\n", terraformVersionPath)
		return nil
	}

	err = os.WriteFile(terraformVersionPath, []byte(updatedVersionFile), 0644)
	if err != nil {
		return err
	}

	beforeVersion := strings.TrimSpace(terraformVersionFile)
	afterVersion := strings.TrimSpace(updatedVersionFile)

	log.Printf("%s updated (%s -> %s)\n", terraformVersionPath, beforeVersion, afterVersion)

	return nil
}

func readFile(file string) (string, error) {
	data, err := os.ReadFile(file)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
