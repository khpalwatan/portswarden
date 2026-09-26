# Changelog

All notable changes to portswarden are documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/),
and this project adheres to [Semantic Versioning](https://semver.org/).

## [Unreleased]

## [0.1.0] - 2026-09-26

### Added
- `list` command — show all listening TCP/UDP ports with process info
- `list -p <port>` — filter to a single port
- `kill <port>` — free a port in one command
- `watch <port>` — auto-kill anything that grabs a port
- `excluded` — show Windows-reserved TCP port ranges
- App icon embedded in the `.exe`
- GitHub Actions workflow for auto-building releases
- MIT license
- README with usage examples and tech stack