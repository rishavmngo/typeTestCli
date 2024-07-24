#!/bin/bash

trap 'echo "Script interrupted. Exiting..."; exit 1;' INT
TARBALL="Typetest-cli.tar.gz"

GREEN='\033[0;32m'
RESET='\033[0m'
print_green() {
  echo -e "${GREEN}$1${RESET}"
}

if [ -f "$TARBALL" ]; then
  rm "$TARBALL"
  print_green "File '$TARBALL' deleted."
fi

print_green "Building go project"
go build -o typeTest

print_green "Creating temperory directory"
mkdir tempBuild

print_green "Copying binary, settings.json and words.json to temp folder"
cp typeTest tempBuild/typeTest
cp settings.json tempBuild/settings.json
cp words.json tempBuild/words.json

print_green "Creating tar file"
tar -czvf $TARBALL tempBuild/

print_green "Deleteing tempBuild folder"
rm -rf tempBuild
