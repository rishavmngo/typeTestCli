#!/bin/bash

TARBALL="Typetest-cli.tar.gz"

if [ -f "$TARBALL" ]; then
  rm "$TARBALL"
  echo "File '$TARBALL' deleted."
fi

echo "Building go project"
go build -o typeTest

echo "Creating temperory directory"
mkdir tempBuild

echo "Copying binary, settings.json and words.json to temp folder"
cp typeTest tempBuild/typeTest
cp settings.json tempBuild/settings.json
cp words.json tempBuild/words.json

echo "Creating tar file"
tar -czvf $TARBALL tempBuild/

echo "Deleteing tempBuild folder"
rm -rf tempBuild
