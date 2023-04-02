#!/usr/bin/env bash
set -euo pipefail

quit() {
	echo "${1}"

	exit 1
}

version="${1}"

if [[ $version =~ v[0-9]\.[0-9]\.[0-9] ]]; then
	echo "Accepting ${version}"
else
	quit "Tag ${version} does not match v1.0.2 format"
fi

echo "Checking for unstaged files"
git status --porcelain
git add .

echo "Checking for uncommitted files"
git diff-index --quiet HEAD -- || quit "git is dirty, ensure all changes are committed"

git push origin || quit "failed to push current commits to origin"

git tag "${version}"
git push origin "${version}"
