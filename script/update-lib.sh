#!/usr/bin/bash
user="mhthrh"
repo="common-lib"
branch="development"

ver="$(curl -s https://api.github.com/repos/$user/$repo/commits/$branch | jq -r '.sha')"

echo "latest commit is: $ver"
go get github.com/mhthrh/common-lib@"$ver"