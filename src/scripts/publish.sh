#!/bin/bash

# Get the current directory of the script
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Execute the publish.sh script
$DIR/semantic-release.sh

npm run format
git add .

# Extract the version from package.json
VERSION=$(grep -o '"version": "[^"]*' package.json | sed 's/"version": "//')

git commit -am "chore: release v$VERSION" --no-verify
git push --no-verify
