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

All flags are optional.

| Flag(s) | Description |
| ------- | ----------- |
| `-f`, `--set-fail` | Set the percentage of a specified failure return code, The format is `rc,percentage`. This flag can be set multiple times. Return codes must be between 1 and 255 and percentages must be between 1 and 100. If multiple failure sets aggregate above 100%, then subsequent failure sets will be ignored that are outside the 100%. This means if we have three sets, `-f 1,40 -f 2,60 -f 3,40` then only `-f 1,40` and `-f 2,60` will be used as return codes. As well, if we have `-f 1,40 -f 2,70` then RC 1 will have a 40% chance and RC 2 will have a 60% chance, ignoring the original 70% value provided. |
| `-s`, `--sleep`    | Set the sleep time in seconds (must be greater or equal to 0) (default: random value between 0-10 seconds) |
| `-V`, `--verbose`  | Enable verbose output |
| `-h`, `--help`     | Help for crash-and-burn |
| `-v`, `--version`  | Version for crash-and-burn |

## Examples

### Single Failure Return Code

- Failure Return Code of `2` with a `30` percent chance of occurring.
- Success Return Code of `0` with a `70` percent chance of occurring.

```sh
crash-and-burn -f 2,30
```

### Multiple Failure Return Codes

- Failure Return Code of `2` with a `30` percent chance of occurring.
- Failure Return Code of `1` with a `10` percent chance of occurring.
- Success Return Code of `0` with a `60` percent chance of occurring.

```sh
crash-and-burn -f 2,30 -f 1,10
```
