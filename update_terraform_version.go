package updater

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/sue445/ghrcooldown"
)

// UpdateTerraformVersionParams represents UpdateTerraformVersion's params
type UpdateTerraformVersionParams struct {
	Src            string
	TargetVersion  string
	UpdaterVersion string
	CooldownDays   int
	CurrentTime    *time.Time
}

// UpdateTerraformVersion updates version in .terraform-version
func UpdateTerraformVersion(params *UpdateTerraformVersionParams) (string, error) {
	if !regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+$`).MatchString(strings.TrimSpace(params.Src)) {
		return params.Src, nil
	}

	if fmt.Sprintf("%s\n", params.TargetVersion) == params.Src {
		return params.Src, nil
	}

	client, err := ghrcooldown.NewClient(&ghrcooldown.ClientParams{
		Token:       os.Getenv("GITHUB_TOKEN"),
		UserAgent:   fmt.Sprintf("terraform-version-updater/%s (+https://github.com/sue445/terraform-version-updater)", params.UpdaterVersion),
		CurrentTime: params.CurrentTime,
	})

	if err != nil {
		return "", errors.WithStack(err)
	}

	cooldown := days(params.CooldownDays)

	if params.TargetVersion == "latest" {
		tagName, err := client.GetLatestTagName(context.Background(), "hashicorp", "terraform", cooldown)
		if err != nil {
			return "", errors.WithStack(err)
		}

		version := strings.TrimPrefix(tagName, "v")
		return fmt.Sprintf("%s\n", version), nil
	}

	hasPassed, err := client.HasCooldownPassed(context.Background(), "hashicorp", "terraform", fmt.Sprintf("v%s", params.TargetVersion), cooldown)
	if err != nil {
		return "", errors.WithStack(err)
	}

	if hasPassed {
		return fmt.Sprintf("%s\n", params.TargetVersion), nil
	}

	return params.Src, nil
}

func days(days int) time.Duration {
	return time.Duration(days) * 24 * time.Hour
}
