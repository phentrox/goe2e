# Go Sequential E2E Test Runner (In Development!)
![Go Report Card](https://goreportcard.com/badge/github.com/phentrox/goseq)
![Lines-of-Code](https://img.shields.io/badge/lines--of--code-458-brightgreen)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/phentrox/goseq/go.yml?branch=main)

<img src=".github/logo.png" alt="Logo" width="400">

## Features
- Fail Fast -> Stops when a test fails
- Sleeps 1 second after every test for cleanups (DB, API, ...)
- Prints individual test time
- Prints total test time

## Requirements
- One folder for every test file
- Every test file needs to end with 'test'
- Every test file only should have one test method! (Keeps sequential order)

## Commands
```sh
# default test dir: testingE2E
goseq

# run with custom test dir
goseq --dir customTesting
goseq -d customTesting
```

## Why?
- Go Unit Tests can run in parallel
- Some E2E needs to run sequentially!
- E.g. DB E2E which runs on the same DB but manipulates the same tables
- This can cause test failures if every DB E2E test assumes a clean table
- Even worse it can succeed sometimes and fail otherwise!
- Solution = Run every test in sequence
- It is slower than parallel execution but provides test stability

## Trademarks
The Gopher mascot in the logo is a trademark of Google. Image created by Renée French. [Source](https://golang.org/doc/gopher/)