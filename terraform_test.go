package main

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetTerraformStableVersions(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://releases.hashicorp.com/terraform/",
		httpmock.NewStringResponder(200, ReadTestData("testdata/terraform-releases.html")))

	want := []string{
		"1.8.5",
		"1.8.4",
		"1.8.3",
		"1.8.2",
		"1.8.1",
		"1.8.0",
		"1.7.5",
		"1.7.4",
		"1.7.3",
		"1.7.2",
		"1.7.1",
		"1.7.0",
		"1.6.6",
		"1.6.5",
		"1.6.4",
		"1.6.3",
		"1.6.2",
		"1.6.1",
		"1.6.0",
		"1.5.7",
		"1.5.6",
		"1.5.5",
		"1.5.4",
		"1.5.3",
		"1.5.2",
		"1.5.1",
		"1.5.0",
		"1.4.7",
		"1.4.6",
		"1.4.5",
		"1.4.4",
		"1.4.3",
		"1.4.2",
		"1.4.1",
		"1.4.0",
		"1.3.10",
		"1.3.9",
		"1.3.8",
		"1.3.7",
		"1.3.6",
		"1.3.5",
		"1.3.4",
		"1.3.3",
		"1.3.2",
		"1.3.1",
		"1.3.0",
		"1.2.9",
		"1.2.8",
		"1.2.7",
		"1.2.6",
		"1.2.5",
		"1.2.4",
		"1.2.3",
		"1.2.2",
		"1.2.1",
		"1.2.0",
		"1.1.9",
		"1.1.8",
		"1.1.7",
		"1.1.6",
		"1.1.5",
		"1.1.4",
		"1.1.3",
		"1.1.2",
		"1.1.1",
		"1.1.0",
		"1.0.11",
		"1.0.10",
		"1.0.9",
		"1.0.8",
		"1.0.7",
		"1.0.6",
		"1.0.5",
		"1.0.4",
		"1.0.3",
		"1.0.2",
		"1.0.1",
		"1.0.0",
		"0.15.5",
		"0.15.4",
		"0.15.3",
		"0.15.2",
		"0.15.1",
		"0.15.0",
		"0.14.11",
		"0.14.10",
		"0.14.9",
		"0.14.8",
		"0.14.7",
		"0.14.6",
		"0.14.5",
		"0.14.4",
		"0.14.3",
		"0.14.2",
		"0.14.1",
		"0.14.0",
		"0.13.7",
		"0.13.6",
		"0.13.5",
		"0.13.4",
		"0.13.3",
		"0.13.2",
		"0.13.1",
		"0.13.0",
		"0.12.31",
		"0.12.30",
		"0.12.29",
		"0.12.28",
		"0.12.27",
		"0.12.26",
		"0.12.25",
		"0.12.24",
		"0.12.23",
		"0.12.22",
		"0.12.21",
		"0.12.20",
		"0.12.19",
		"0.12.18",
		"0.12.17",
		"0.12.16",
		"0.12.15",
		"0.12.14",
		"0.12.13",
		"0.12.12",
		"0.12.11",
		"0.12.10",
		"0.12.9",
		"0.12.8",
		"0.12.7",
		"0.12.6",
		"0.12.5",
		"0.12.4",
		"0.12.3",
		"0.12.2",
		"0.12.1",
		"0.12.0",
		"0.11.15",
		"0.11.14",
		"0.11.13",
		"0.11.12",
		"0.11.11",
		"0.11.10",
		"0.11.9",
		"0.11.8",
		"0.11.7",
		"0.11.6",
		"0.11.5",
		"0.11.4",
		"0.11.3",
		"0.11.2",
		"0.11.1",
		"0.11.0",
		"0.10.8",
		"0.10.7",
		"0.10.6",
		"0.10.5",
		"0.10.4",
		"0.10.3",
		"0.10.2",
		"0.10.1",
		"0.10.0",
		"0.9.11",
		"0.9.10",
		"0.9.9",
		"0.9.8",
		"0.9.7",
		"0.9.6",
		"0.9.5",
		"0.9.4",
		"0.9.3",
		"0.9.2",
		"0.9.1",
		"0.9.0",
		"0.8.8",
		"0.8.7",
		"0.8.6",
		"0.8.5",
		"0.8.4",
		"0.8.3",
		"0.8.2",
		"0.8.1",
		"0.8.0",
		"0.7.13",
		"0.7.12",
		"0.7.11",
		"0.7.10",
		"0.7.9",
		"0.7.8",
		"0.7.7",
		"0.7.6",
		"0.7.5",
		"0.7.4",
		"0.7.3",
		"0.7.2",
		"0.7.1",
		"0.7.0",
		"0.6.16",
		"0.6.15",
		"0.6.14",
		"0.6.13",
		"0.6.12",
		"0.6.11",
		"0.6.10",
		"0.6.9",
		"0.6.8",
		"0.6.7",
		"0.6.6",
		"0.6.5",
		"0.6.4",
		"0.6.3",
		"0.6.2",
		"0.6.1",
		"0.6.0",
		"0.5.3",
		"0.5.1",
		"0.5.0",
		"0.4.2",
		"0.4.1",
		"0.4.0",
		"0.3.7",
		"0.3.6",
		"0.3.5",
		"0.3.1",
		"0.3.0",
		"0.2.2",
		"0.2.1",
		"0.2.0",
		"0.1.1",
		"0.1.0",
	}

	got, err := GetTerraformStableVersions()
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

// ReadTestData returns testdata
func ReadTestData(filename string) string {
	buf, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(buf)
}
