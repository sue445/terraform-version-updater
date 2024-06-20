package updater

import (
	"os"
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

	result := UpdateTerraformVersion(&UpdateTerraformVersionParams{
		Src:           terraformVersionFile,
		TargetVersion: targetVersion,
		Versions:      versions,
	})

	if result == terraformVersionFile {
		// version isn't updated
		return nil
	}

	err = os.WriteFile(terraformVersionPath, []byte(result), 0644)
	if err != nil {
		return err
	}

	return nil
}

func readFile(file string) (string, error) {
	data, err := os.ReadFile(file)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
