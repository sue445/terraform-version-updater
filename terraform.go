package updater

import (
	"github.com/cockroachdb/errors"
	"io"
	"net/http"
	"regexp"
)

// GetTerraformStableVersions returns all stable versions from https://releases.hashicorp.com/terraform/
func GetTerraformStableVersions() ([]string, error) {
	html, err := getTerraformReleasesHTML()
	if err != nil {
		return []string{}, errors.WithStack(err)
	}

	matches := regexp.MustCompile(`<a href="/terraform/([0-9]+\.[0-9]+\.[0-9]+)/">`).FindAllStringSubmatch(html, -1)

	var versions []string
	for _, match := range matches {
		if len(match) > 1 {
			versions = append(versions, match[1])
		}
	}

	return versions, nil
}

func getTerraformReleasesHTML() (string, error) {
	// c.f. https://github.com/tfutils/tfenv/blob/master/libexec/tfenv-list-remote
	res, err := http.Get("https://releases.hashicorp.com/terraform/")
	if err != nil {
		return "", errors.WithStack(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", errors.New(res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return string(body), nil
}
