package updater_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/sue445/terraform-version-updater"
	"testing"
)

func TestUpdateTerraformVersion(t *testing.T) {
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
				Versions: []string{
					"1.8.5",
					"1.8.4",
					"1.8.3",
					"1.8.2",
					"1.8.1",
					"1.8.0",
				},
			},
			want: "1.8.5\n",
		},
		{
			name: "Update to specified version",
			args: &updater.UpdateTerraformVersionParams{
				Src:           "1.8.0\n",
				TargetVersion: "1.8.4",
				Versions: []string{
					"1.8.5",
					"1.8.4",
					"1.8.3",
					"1.8.2",
					"1.8.1",
					"1.8.0",
				},
			},
			want: "1.8.4\n",
		},
		{
			name: "Doesn't Updated",
			args: &updater.UpdateTerraformVersionParams{
				Src:           "1.8.0\n",
				TargetVersion: "1.8.0",
				Versions: []string{
					"1.8.5",
					"1.8.4",
					"1.8.3",
					"1.8.2",
					"1.8.1",
					"1.8.0",
				},
			},
			want: "1.8.0\n",
		},
		{
			name: ".terraform-version is invalid",
			args: &updater.UpdateTerraformVersionParams{
				Src:           "abcdef\n",
				TargetVersion: "1.8.0",
				Versions: []string{
					"1.8.5",
					"1.8.4",
					"1.8.3",
					"1.8.2",
					"1.8.1",
					"1.8.0",
				},
			},
			want: "abcdef\n",
		},
		{
			name: "targetVersion is unknown",
			args: &updater.UpdateTerraformVersionParams{
				Src:           "1.8.0\n",
				TargetVersion: "1.7.5",
				Versions: []string{
					"1.8.5",
					"1.8.4",
					"1.8.3",
					"1.8.2",
					"1.8.1",
					"1.8.0",
				},
			},
			want: "1.8.0\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := updater.UpdateTerraformVersion(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}
