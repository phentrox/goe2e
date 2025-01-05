# Go Sequential E2E Testing (In Development!)

## Readme
- One folder for every test file
- Every test file needs to end with 'test'
- Every test file only should have one test method! (Keeps sequential order)

## Convention over Configuration
- Stops when a test fails
- Sleeps 1 second after every test for cleanups (DB, ...)

## Why?
- Go Unit Tests can run in parallel
- Some E2E needs to run sequentially!
- E.g. DB E2E which runs on the same DB but manipulates the same tables
- This can cause test failures if every DB E2E test assumes a clean table
- Even worse it can succeed sometimes and fail otherwise!
- Solution = Run every test in sequence
- Is slower than parallel execution but provides test stability
