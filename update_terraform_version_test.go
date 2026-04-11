package updater_test

import (
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/sue445/terraform-version-updater"
)

func TestUpdateTerraformVersion_Success(t *testing.T) {
	httpmock.Activate(t)
	httpmock.RegisterResponder(
		"GET",
		"https://api.github.com/repos/hashicorp/terraform/releases?per_page=10",
		httpmock.NewStringResponder(200, readFile(t, "testdata/terraform-releases.json")),
	)
	httpmock.RegisterResponder(
		"GET",
		"https://api.github.com/repos/hashicorp/terraform/releases/tags/v1.14.8",
		httpmock.NewStringResponder(200, readFile(t, "testdata/terraform-tags-v1.14.8.json")),
	)
	httpmock.RegisterResponder(
		"GET",
		"https://api.github.com/repos/hashicorp/terraform/releases/tags/v1.14.7",
		httpmock.NewStringResponder(200, readFile(t, "testdata/terraform-tags-v1.14.7.json")),
	)

	tests := []struct {
		name string
		args *updater.UpdateTerraformVersionParams
		want string
	}{
		{
			name: "Update to latest version",
			args: &updater.UpdateTerraformVersionParams{
				Src:           "1.8.0\n",
				TargetVersion: "latest",
				CurrentTime:   new(time.Date(2026, 3, 26, 0, 0, 0, 0, time.UTC)),
				CooldownDays:  7,
			},
			want: "1.14.7\n",
		},
		{
			name: "Update to specified version",
			args: &updater.UpdateTerraformVersionParams{
				Src:           "1.8.0\n",
				TargetVersion: "1.14.8",
				CurrentTime:   new(time.Date(2026, 3, 26, 0, 0, 0, 0, time.UTC)),
				CooldownDays:  0,
			},
			want: "1.14.8\n",
		},
		{
			name: "Doesn't Updated",
			args: &updater.UpdateTerraformVersionParams{
				Src:           "1.8.0\n",
				TargetVersion: "1.8.0",
				CurrentTime:   new(time.Date(2026, 3, 26, 0, 0, 0, 0, time.UTC)),
				CooldownDays:  7,
			},
			want: "1.8.0\n",
		},
		{
			name: ".terraform-version is invalid",
			args: &updater.UpdateTerraformVersionParams{
				Src:           "abcdef\n",
				TargetVersion: "1.8.0",
				CurrentTime:   new(time.Date(2026, 3, 26, 0, 0, 0, 0, time.UTC)),
				CooldownDays:  7,
			},
			want: "abcdef\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := updater.UpdateTerraformVersion(tt.args)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestUpdateTerraformVersion_TargetVersionIsUnknown(t *testing.T) {
	httpmock.Activate(t)

	httpmock.RegisterResponder(
		"GET",
		"https://api.github.com/repos/hashicorp/terraform/releases/tags/v1.14.0-unknown",
		httpmock.NewStringResponder(404, readFile(t, "testdata/terraform-tags-v1.14.0-unknown.json")),
	)

	_, err := updater.UpdateTerraformVersion(&updater.UpdateTerraformVersionParams{
		Src:           "1.8.0\n",
		TargetVersion: "1.14.0-unknown",
		CurrentTime:   new(time.Date(2026, 3, 26, 0, 0, 0, 0, time.UTC)),
		CooldownDays:  7,
	})

	if assert.Error(t, err) {
		assert.ErrorContains(t, err, "https://api.github.com/repos/hashicorp/terraform/releases/tags/v1.14.0-unknown")
		assert.ErrorContains(t, err, "404 Not Found")
	}
}
