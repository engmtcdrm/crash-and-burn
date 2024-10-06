# crash-and-burn

<a href="https://github.com/engmtcdrm/crash-and-burn/actions"><img src="https://github.com/engmtcdrm/crash-and-burn/actions/workflows/build.yml/badge.svg?branch=main" alt="Build Status"></a>

A simple utility for randomly generating error, warning, and success return codes.

This utility program is meant to be used for development of other programs that may call subprocesses to do something.

## Usage

```sh
crash-and-burn
```

### With Flags

```sh
crash-and-burn -r
```

## Flags

All options are optional.

| Flag(s) | Description |
| ------- | ----------- |
| `--err-pct`       | Set the error percentage (default: `2`) |
| `--err-rc`        | Set the error return code (default: `99`) |
| `--warn-pct`      | Set the warning percentage (default: `30`) |
| `--warn-rc`       | Set the warning return code (default: `1`) |
| `-s`, `--sleep`   | Set the sleep time in seconds (default: `random`) |
| `-d`, `--debug`   | Enable debug mode |
| `-h`, `--help`    | help for crash-and-burn |
| `-v`, `--version` | version for crash-and-burn |
