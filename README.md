# terraform-version-updater
Automatically update [`.terraform-version`](https://github.com/tfutils/tfenv)

[![Latest Version](https://img.shields.io/github/v/release/sue445/terraform-version-updater)](https://github.com/sue445/terraform-version-updater/releases)
[![test](https://github.com/sue445/terraform-version-updater/actions/workflows/test.yml/badge.svg)](https://github.com/sue445/terraform-version-updater/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/github/sue445/terraform-version-updater/badge.svg)](https://coveralls.io/github/sue445/terraform-version-updater)
[![Go Report Card](https://goreportcard.com/badge/github.com/sue445/terraform-version-updater)](https://goreportcard.com/report/github.com/sue445/terraform-version-updater)
[![Go Reference](https://pkg.go.dev/badge/github.com/sue445/terraform-version-updater.svg)](https://pkg.go.dev/github.com/sue445/terraform-version-updater)

## Install
Download latest binary from https://github.com/sue445/terraform-version-updater/releases

## Build
```bash
go install github.com/sue445/terraform-version-updater@latest
```

## Example
```bash
cd /path/to/terraform-repo

# Update terraform to latest version
terraform-version-updater

# Update terraform to latest version (dry-run)
terraform-version-updater --dry-run

# Update terraform to specified version
terraform-version-updater --target 1.8.5

# Update terraform to latest version with specified .terraform-version file
terraform-version-updater --file /path/to/.terraform-version

# Show terraform-version-updater's version
terraform-version-updater --version
```

## Usage
```bash
$ terraform-version-updater --help
Usage of terraform-version-updater:
  -d, --dry-run         Whether dry-run
  -f, --file string     Path to .terraform-version file (default ".terraform-version")
  -h, --help            Whether show help
  -t, --target string   Version to be updated (default "latest")
  -v, --version         Whether showing version
```

## vs [tfupdate](https://github.com/minamijoyo/tfupdate)
* _terraform-version-updater_ supports `.terraform-version`
* _tfupdate_ supports `required_version` in `*.tf`

## GitHub Actions Example
Put the following yaml to your repository (e.g. `.github/workflows/terraform-version-updater.yml`)

e.g.

```yml
name: Upgrade Terraform to latest version

on:
  schedule:
    - cron: "0 0 1 * *" # Run monthly
  workflow_dispatch: # Run manually

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4

      - name: Get latest release info
        id: get_latest_release
        uses: octokit/request-action@v2.x
        with:
          route: GET /repos/sue445/terraform-version-updater/releases/latest
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

      - name: Add LATEST_VERSION to GITHUB_ENV
        run: echo "LATEST_VERSION=${{ fromJson(steps.get_latest_release.outputs.data).tag_name }}" >> $GITHUB_ENV

      - name: Download and install terraform-version-updater
        run: |
          wget https://github.com/sue445/terraform-version-updater/releases/download/${LATEST_VERSION}/terraform-version-updater_Linux_x86_64.tar.gz
          tar -zxvf terraform-version-updater_Linux_x86_64.tar.gz
          mv terraform-version-updater /usr/local/bin
        working-directory: /tmp

      - name: Add BEFORE_TERRAFORM_VERSION to GITHUB_ENV
        run: echo "BEFORE_TERRAFORM_VERSION=$(cat .terraform-version)" >> $GITHUB_ENV

      - name: Run terraform-version-updater
        run: terraform-version-updater

      - name: Add AFTER_TERRAFORM_VERSION to GITHUB_ENV
        run: echo "AFTER_TERRAFORM_VERSION=$(cat .terraform-version)" >> $GITHUB_ENV

      - name: Create Terraform version up PullRequest
        uses: peter-evans/create-pull-request@v6
        with:
          token:          ${{ secrets.GITHUB_TOKEN }}
          title:          "Bump Terraform from ${{ env.BEFORE_TERRAFORM_VERSION }} to ${{ env.AFTER_TERRAFORM_VERSION }}"
          commit-message: "Bump Terraform from ${{ env.BEFORE_TERRAFORM_VERSION }} to ${{ env.AFTER_TERRAFORM_VERSION }}"
          labels:         "terraform-version-updater"
          branch:         "terraform-version-updater/terraform_${{ env.AFTER_TERRAFORM_VERSION }}"
          body: |
            Bumps [Terraform](https://github.com/hashicorp/terraform) from ${{ env.BEFORE_TERRAFORM_VERSION }} to ${{ env.AFTER_TERRAFORM_VERSION }}

            * Release: https://github.com/hashicorp/terraform/releases/tag/v${{ env.AFTER_TERRAFORM_VERSION }}
            * See full diff in [compare view](https://github.com/hashicorp/terraform/compare/v${{ env.BEFORE_TERRAFORM_VERSION }}...v${{ env.AFTER_TERRAFORM_VERSION }})
```

When using this workflow, it is recommended to use `.terraform-version` in `hashicorp/setup-terraform` as follows

e.g.

```yml
- name: Set variables
  run: |
    echo "TERRAFORM_VERSION=$(cat .terraform-version)" >> $GITHUB_ENV

- uses: hashicorp/setup-terraform@v3
  with:
    terraform_version: ${{ env.TERRAFORM_VERSION }}
```

### Known problem
GitHub Actions don't allow recursive builds.

So a Pull Request created using `secrets.GITHUB_TOKEN` will not execute build.

c.f. https://docs.github.com/en/actions/security-guides/automatic-token-authentication#using-the-github_token-in-a-workflow

Therefore, the use of App Token by [`actions/create-github-app-token`](https://github.com/marketplace/actions/create-github-app-token) is **strongly recommended**.

e.g.

```yml
- uses: actions/create-github-app-token@v1
  id: app-token
  with:
    app-id: ${{ secrets.APP_ID }}
    private-key: ${{ secrets.PRIVATE_KEY }}

- name: Create Terraform version up PullRequest
  uses: peter-evans/create-pull-request@v6
  with:
    # Use steps.app-token.outputs.token instead of secrets.GITHUB_TOKEN
    token: ${{ steps.app-token.outputs.token }}
```
