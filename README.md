# :fire: crash-and-burn

[![Go](https://github.com/engmtcdrm/crash-and-burn/actions/workflows/build.yml/badge.svg)](https://github.com/engmtcdrm/crash-and-burn/actions/workflows/build.yml)
[![Release](https://img.shields.io/github/v/release/engmtcdrm/crash-and-burn.svg?label=Latest%20Release)](https://github.com/engmtcdrm/minno/releases/latest)

A simple utility for randomly generating success and failure return codes.

This utility program is meant to be used when developing other programs that call subprocesses.

## :zap: Usage

```sh
crash-and-burn [flags]
```

## Flags

All options are optional.

| Flag(s) | Description |
| ------- | ----------- |
| `-f`, `--set-fail` | set the percentage of a specified failure return code, The format is rc,percentage. Can be set multiple times. Return codes must be between 1 and 255 and percentages must be between 1 and 100. |
| `-s`, `--sleep`    | set the sleep time in seconds (must be greater or equal to 0) (default: random value between 0-10 seconds) |
| `-V`, `--verbose`  | enable verbose output |
| `-h`, `--help`     | help for crash-and-burn |
| `-v`, `--version`        | version for crash-and-burn |

## Examples

### Single Failure RC

- Return Code of `2` and percentage of `30`

```sh
crash-and-burn -f 2,30
```

### Multiple Failure RCs

- Return Code of `2` and percentage of `30`
- Return Code of `1` and percentage of `10`

```sh
crash-and-burn -f 2,30 -f 1,10
```
