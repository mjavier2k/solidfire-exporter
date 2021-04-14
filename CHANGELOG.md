# Changelog
<!-- https://keepachangelog.com/en/1.0.0/

Types of Changes:
Added - for new features.
Changed - for changes in existing functionality.
Deprecated - for soon-to-be removed features.
Removed - for now removed features.
Fixed - for any bug fixes.
Security - in case of vulnerabilities.

Example Format Below:

## [0.0.7] - 2015-02-16
### Added
- Link, and make it obvious that date format is ISO 8601.

### Changed
- Clarified the section on "Is there a standard change log format?".

### Fixed
- Fix Markdown links to tag comparison URL with footnote-style links.
-->
## [Unreleased]

### Added
- Added CHANGELOG
### Changed
### Deprecated
### Removed
### Fixed
### Security

## [0.5.2] - 2021-03-11

### Added
- Added license [#32](https://github.com/mjavier2k/solidfire-exporter/pull/32)
- Added official dashboards [#39](https://github.com/mjavier2k/solidfire-exporter/pull/39)
- Added Arm64 Artifact [#46](https://github.com/mjavier2k/solidfire-exporter/pull/46)
- Officially supports ElementOS 12 [#48](https://github.com/mjavier2k/solidfire-exporter/pull/48)

### Changes
- Changed scrape_success metric to solidfire_up [#29](https://github.com/mjavier2k/solidfire-exporter/pull/29)
- Added option to specify config.yaml (NOTE: Environment variable names changed) [#30](https://github.com/mjavier2k/solidfire-exporter/pull/30)

### Fixed
- Protect against bad URI causing crash [#47](https://github.com/mjavier2k/solidfire-exporter/pull/47)
- Fix efficiency factor calculation [#49](https://github.com/mjavier2k/solidfire-exporter/pull/49)
