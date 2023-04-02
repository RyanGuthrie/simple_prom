#!/usr/bin/env zsh
set -euo pipefail

autoload colors; colors

quit() {
	echo -e "$fg[red]${1} ${reset_color}"

	exit 1
}

SCRIPT_DIR=$( cd -- "$( dirname -- "${ZSH_ARGZERO}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR

if (( "${#}" != 1 )); then
	echo "Latest tag is:"
	echo -n "  "
	git tag | sort -V | tail -n1
	echo
	quit "Usage: ${0} v1.0.10"
fi

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
echo

git push origin || quit "failed to push current commits to origin"

git tag "${version}"
git push origin "${version}"
