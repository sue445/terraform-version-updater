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
