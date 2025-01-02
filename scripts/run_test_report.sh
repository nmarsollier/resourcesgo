#!/bin/bash
# Remove any existing coverage files
rm -rf ./coverage

# Create a directory to store individual coverage profiles
mkdir -p coverage

# Find all packages and run tests with coverage
go test -coverprofile=coverage/coverage.out ./...  > /dev/null 2>&1

go tool cover -func=./coverage/coverage.out | grep "total:" |  tr -d '[:space:]' | sed 's/(statements)//g' | sed 's/%*$//'

# Generate an HTML report
go tool cover -html=./coverage/coverage.out -o coverage/coverage.html

if [ "$1" != "quiet" ]; then
  nohup open coverage/coverage.html > /dev/null 2>&1&
fi
 