 Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## 0.0.2 - 2025-03-04

### Added

- Interfaces for `UserIDValidatorPlugin`s.
- `Initializable` interface for stateful plugins.

### Changed

- **Breaking Change** `ctx context.Context` has been added as a first argument  for the `ValidateKey` method in the `AppKeyValidationPlugin` interface.

## 0.0.1 - 2025-02-26

### Added

- `UserGetterPlugin` interface
- `AppKeyGetterPlugin` interface
- `AppKeyValidationPlugin` interface