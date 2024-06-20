package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

// UpdateTerraformVersion updates version in .terraform-version
func UpdateTerraformVersion(src string, targetVersion string, versions []string) string {
	if !regexp.MustCompile("^[0-9]+\\.[0-9]+\\.[0-9]+$").MatchString(strings.TrimSpace(src)) {
		return src
	}

	if targetVersion == "latest" {
		return fmt.Sprintf("%s\n", versions[0])
	}

	if slices.Contains(versions, targetVersion) {
		return fmt.Sprintf("%s\n", targetVersion)
	}

	return src
}
