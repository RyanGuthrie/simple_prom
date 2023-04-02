#!/usr/bin/env zsh
set -euo pipefail

autoload colors; colors

quit() {
	echo -e "$fg[red]${1}${reset_color}"

	exit 1
}

info() {
	echo -e "$fg[blue]${1}${reset_color}"
}

SCRIPT_DIR=$( cd -- "$( dirname -- "${ZSH_ARGZERO}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR

if (( "${#}" != 1 )); then
	info "Latest tag is:"

	echo -n "  "
	git tag | sort -V | tail -n1

	echo
	info "Usage: ${0} v1.0.10"

	exit 1
fi

version="${1}"

if [[ $version =~ v[0-9]\.[0-9]\.[0-9] ]]; then
	info "Accepting tag ${version}"
else
	quit "Tag ${version} does not match v1.0.2 format"
fi

info "Checking for unstaged files"
git status --porcelain
git add .

info "Checking for uncommitted files"
git diff-index --quiet HEAD -- || quit "git is dirty, ensure all changes are committed"

git push origin || quit "failed to push current commits to origin"

info "Tagging as ${version}"
git tag "${version}"

info "Pushing tag ${version} to origin"
git push origin "${version}"

info "Completed tag successfully"
