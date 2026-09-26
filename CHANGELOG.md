# Changelog

All notable changes to portswarden are documented here.

## [1.0.0] - 2026-09-26

First stable release.

### Added
- `list` command now supports `--json` output
- `list` command now supports `--sort` (port, pid, process)
- `kill` command now shows a confirmation prompt (skip via `--yes`)
- `doctor` command for environment diagnostics
- `--version` flag

### Changed
- `list` now shows `<protected>` for restricted system processes instead of blank cells
- Versioning moved to stable SemVer (1.x.x)

## [0.1.0] - 2026-09-26

Initial alpha release.

### Added
- `list` command — show all listening TCP/UDP ports with process info
- `list -p <port>` — filter to a single port
- `kill <port>` — free a port in one command
- `watch <port>` — auto-kill anything that grabs a port
- `excluded` — show Windows-reserved TCP port ranges
- App icon embedded in the `.exe`
- GitHub Actions workflow for auto-releases
- MIT license