package updater_test

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/sue445/terraform-version-updater"
	"os"
	"path/filepath"
	"testing"
)

func TestExecute(t *testing.T) {
	httpmock.Activate(t)

	httpmock.RegisterResponder("GET", "https://releases.hashicorp.com/terraform/",
		httpmock.NewStringResponder(200, readFile(t, "testdata/terraform-releases.html")))

	type args struct {
		targetVersion        string
		terraformVersionFile string
		isDryRun             bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Update to latest",
			args: args{
				targetVersion:        "latest",
				terraformVersionFile: "1.8.0\n",
				isDryRun:             false,
			},
			want: "1.8.5\n",
		},
		{
			name: "Update to latest (dry-run)",
			args: args{
				targetVersion:        "latest",
				terraformVersionFile: "1.8.0\n",
				isDryRun:             true,
			},
			want: "1.8.0\n",
		},
		{
			name: "Doesn't updated",
			args: args{
				targetVersion:        "1.8.5",
				terraformVersionFile: "1.8.5\n",
				isDryRun:             false,
			},
			want: "1.8.5\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()

			terraformVersionPath := filepath.Join(dir, ".terraform-version")
			createFile(t, terraformVersionPath, tt.args.terraformVersionFile)

			u := updater.NewUpdater(tt.args.isDryRun)
			err := u.Execute(tt.args.targetVersion, terraformVersionPath)
			if assert.NoError(t, err) {
				got := readFile(t, terraformVersionPath)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func createFile(t *testing.T, file string, content string) {
	err := os.WriteFile(file, []byte(content), 0644)

	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
}

func readFile(t *testing.T, file string) string {
	data, err := os.ReadFile(file)

	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	return string(data)
}
