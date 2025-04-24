#!/bin/sh
# scripts/completions.sh
set -e
rm -rf completions
mkdir completions
# TODO: replace your-cli with your binary name
for sh in bash zsh fish; do
	go run main.go completion "$sh" >"completions/sshabu.$sh"
done