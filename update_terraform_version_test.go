package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateTerraformVersion(t *testing.T) {
	type args struct {
		src           string
		targetVersion string
		versions      []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Update to latest version",
			args: args{
				src:           "1.8.0\n",
				targetVersion: "latest",
				versions: []string{
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
			args: args{
				src:           "1.8.0\n",
				targetVersion: "1.8.4",
				versions: []string{
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
			args: args{
				src:           "1.8.0\n",
				targetVersion: "1.8.0",
				versions: []string{
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
			args: args{
				src:           "abcdef\n",
				targetVersion: "1.8.0",
				versions: []string{
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
			args: args{
				src:           "1.8.0\n",
				targetVersion: "1.7.5",
				versions: []string{
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
			got := UpdateTerraformVersion(tt.args.src, tt.args.targetVersion, tt.args.versions)
			assert.Equal(t, tt.want, got)
		})
	}
}
