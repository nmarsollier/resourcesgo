#!/bin/bash
# Remove any existing coverage files
rm -rf ./coverage

# Create a directory to store individual coverage profiles
mkdir -p coverage

# Find all packages and run tests with coverage
go test -coverprofile=coverage/coverage.out.tmp ./...  > /dev/null 2>&1

cat coverage/coverage.out.tmp \
 | grep -v "/internal/graph/" \
 | grep -v "/resourcesgo/docs/" \
 | grep -v "/resourcesgo/internal/tools/db/" \
 > coverage/coverage.out

go tool cover -func=./coverage/coverage.out | grep "total:" |  tr -d '[:space:]' | sed 's/(statements)//g' | sed 's/%*$//'

# Generate an HTML report
go tool cover -html=./coverage/coverage.out -o coverage/coverage.html

if [ "$1" != "quiet" ]; then
  nohup open coverage/coverage.html > /dev/null 2>&1&
fi
 