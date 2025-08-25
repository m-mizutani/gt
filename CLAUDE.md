# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`gt` is a Go test library that leverages Go generics to provide type-safe testing with IDE support and compile-time type checking. The library provides wrapper types around Go's standard `testing.TB` to create strongly-typed test assertions.

## Restrictions and Rules

### Directory

- When you are mentioned about `tmp` directory, you SHOULD NOT see `/tmp`. You need to check `./tmp` directory from root of the repository.

### Exposure policy

In principle, do not trust developers who use this library from outside

- Do not export unnecessary methods, structs, and variables
- Assume that exposed items will be changed. Never expose fields that would be problematic if changed
- Use `export_test.go` for items that need to be exposed for testing purposes
- **Exception**: Domain models (`pkg/domain/model/*`) can have exported fields as they represent data structures

### Data Validation and Normalization Policy

- **NEVER relax validation to accept invalid data formats**
- **Always normalize data at the boundary (input/output) layers**
- When reading from external sources (DB, API), normalize data to the correct format immediately
- When writing to external sources, ensure data is in the correct format
- Validation should be strict - accepting invalid formats leads to bugs and technical debt
- Example: LLM providers must always be lowercase ("openai", "claude", "gemini"), not mixed case
- If legacy data exists in wrong format, normalize it when reading, don't relax validation

### Check

When making changes, before finishing the task, always:
- Run `go vet ./...`, `go fmt ./...` to format the code
- Run `golangci-lint run ./...` to check lint error
- Run `gosec -exclude-generated -quiet ./...` to check security issue
- Run `go test ./...` to check side effect
- **For GraphQL changes: Run `task graphql` and verify no compilation errors**
- **For GraphQL changes: Check frontend GraphQL queries are updated accordingly**

### Language

All comment and character literal in source code must be in English

### Testing

- Test files should have `package {name}_test`. Do not use same package name
- **🚨 CRITICAL RULE: Test MUST be included in same name test file. (e.g. test for `abc.go` must be in `abc_test.go`) 🚨**

#### Repository Testing Strategy
- **🚨 CRITICAL: Repository tests MUST be placed in `pkg/repository/database/` directory with common test suites**
- Create shared test functions that verify identical behavior across all repository implementations (Firestore, Memory, etc.)
- Each repository implementation must pass the exact same test suite to ensure behavioral consistency
- Use a common test interface pattern to test all implementations uniformly
- This ensures that switching between repository implementations (e.g., Memory for testing, Firestore for production) maintains identical behavior

## Architecture

### Core Components

- **`generic.go`**: Base types and functions including `baseTest[T]`, `Equal`, `NotEqual`, `Nil`, `NotNil`
- **`interfaces.go`**: Internal testing interfaces with fail-fast behavior
- **Type-specific test wrappers**:
  - `value.go`: `ValueTest[T]` - basic value comparison methods
  - `array.go`: `ArrayTest[T]` - slice/array specific methods (Has, Contains, Length, etc.)
  - `map.go`: `MapTest[K,V]` - map specific methods (HasKey, HasValue, etc.)
  - `number.go`: `NumberTest[T]` - numeric comparison methods (Greater, Less, etc.)
  - `string.go`: `StringTest` - string specific methods
  - `bool.go`: `BoolTest` - boolean specific methods
  - `error.go`: `ErrorTest` - error handling methods
  - `file.go`: `FileTest` - file operation methods
  - `cast.go`: `CastTest[T]` - type casting with nil checks

### Key Design Patterns

- **Generic Type Safety**: Each test type is parameterized with Go generics to ensure compile-time type checking
- **Fluent Interface**: Methods return the same type to allow method chaining
- **Required Pattern**: All test types have a `Required()` method that fails fast on previous errors
- **Sugar Syntax**: Short function names (e.g., `V()` for `Value()`, `A()` for `Array()`)

## Development Commands

### Testing
```bash
go test -v ./...              # Run all tests
go test -v ./... -run TestName # Run specific test
```

### Linting
```bash
go vet ./...                  # Static analysis (used in CI)
```

### Build
```bash
go build ./...                # Build all packages
```

## Testing Conventions

- Test files follow `*_test.go` naming convention
- Each core module has corresponding test file (e.g., `value.go` -> `value_test.go`)
- Tests use the library's own testing functions for meta-testing
- Mock testing is handled in `mock_test.go`

## Key Utilities

- **`util.go`**: Contains comparison functions (`EvalCompare`, `EvalIsNil`) and diff generation (`Diff`)
- **`return.go`**: Helper for testing function return values
- **Error handling**: Custom `errorWithFail` wrapper that calls `FailNow()` on error

## Code Style Notes

- Uses Go 1.19+ with generics
- Single dependency: `github.com/google/go-cmp` for deep comparison
- Helper calls (`t.Helper()`) are used consistently to provide accurate stack traces
- Methods use receiver name `x` consistently across all test types