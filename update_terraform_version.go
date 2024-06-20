package updater

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

// UpdateTerraformVersionParams represents UpdateTerraformVersion's params
type UpdateTerraformVersionParams struct {
	Src           string
	TargetVersion string
	Versions      []string
}

// UpdateTerraformVersion updates version in .terraform-version
func UpdateTerraformVersion(params *UpdateTerraformVersionParams) string {
	if !regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+$`).MatchString(strings.TrimSpace(params.Src)) {
		return params.Src
	}

	if params.TargetVersion == "latest" {
		return fmt.Sprintf("%s\n", params.Versions[0])
	}

	if slices.Contains(params.Versions, params.TargetVersion) {
		return fmt.Sprintf("%s\n", params.TargetVersion)
	}

	return params.Src
}
