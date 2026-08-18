#!/usr/bin/env bash
#
# Reports every function with 0% test coverage. Never fails the build:
# this is a warning-only signal for `make codecov` / `make dod`.
#
# internal/fixture is excluded: it's test-support code (builders reused
# by other packages' tests, see .github/skills/unit-tests.md) with no
# test file of its own, so it always shows 0% and isn't a real gap.

YELLOW='\033[0;33m'
CYAN='\033[0;36m'
NC='\033[0m'

GO_PATHS=$(go list -f '{{ .Dir }}' ./... | grep -E -v 'docs|cmd|mocks|internal/fixture')

COVERAGE_FILE=$(mktemp)
trap 'rm -f "$COVERAGE_FILE"' EXIT

go test -coverprofile="$COVERAGE_FILE" $GO_PATHS > /dev/null 2>&1

uncovered=$(go tool cover -func="$COVERAGE_FILE" 2>/dev/null | awk -F'\t' '{gsub(/ /,"",$NF)} $NF=="0.0%" && $1!="total:"')

if [ -z "$uncovered" ]; then
  printf "${CYAN}codecov:${NC} no functions with 0%% coverage\n"
  exit 0
fi

printf "${YELLOW}WARNING:${NC} the following functions have no test coverage:\n"
printf "${YELLOW}%s${NC}\n" "$uncovered"
exit 0
