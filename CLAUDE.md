# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A collection of standalone Go examples demonstrating various Go language features and patterns. Each subdirectory is an independent program with its own `main.go` (package main). There is no shared library or unified build target.

Module path: `github.com/water25234/golang-practice`

## Build & Run Commands

```bash
# Run any example (each directory is a standalone program)
go run ./<directory>/...

# Example
go run ./channelSignal/main.go

# Run tests (only bytesconv and goroutineAndScheduler have tests)
go test ./bytesconv/...
go test ./goroutineAndScheduler/...

# Run all tests
go test ./...

# Run a single test function
go test ./bytesconv/... -run TestName
```

## Architecture

There is no overarching architecture — each subdirectory is a self-contained example:

- **concurrency/case*-design/** — progressive concurrency patterns; `case2-design` implements a worker pool with `Job`, `Worker`, and `WorkerPool` abstractions across multiple files
- **restful/case2/** — HTTP API example with a separate `api/` package
- Everything else is a single `main.go` demonstrating one concept (channels, interfaces, sync primitives, etc.)

## CI

Drone CI is configured (`.drone.yml`) using `golang:1.13-alpine`. The pipeline runs `go build`, `go test`, then cleans up. Note: the go.mod has no Go version directive, and the Drone image is dated — local development may use a newer Go version.
