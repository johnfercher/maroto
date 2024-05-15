#!/usr/bin/env bash

RED='\033[0;31m'
CYAN='\033[0;36m'
NC='\033[0m'

mockOccurrences="`grep -n --include \*.go --exclude-dir=vendor -R mocks.`"
mocksNotInImports="`echo "$mockOccurrences" | grep -v "/mocks"`"
mocksCreatedWithoutNew="`echo "$mocksNotInImports" | grep -v "New"`"

if [ -z "$mocksCreatedWithoutNew" ];
then
  exit 0
else
  printf "${RED}ERROR:${NC} there are mocks that doesn't follow the constructor pattern ${CYAN}m := mocks.NewConstructor(t)${NC}:\n"
  printf "${RED}$mocksCreatedWithoutNew"
  exit 1
fi