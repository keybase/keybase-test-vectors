#!/bin/sh
# Example Usage:
#   ./compile.sh foo.iced bar.iced

for i in $*; do
    input=$i
    output=$(basename -s .iced "$i").json
    echo "$input                   -> ../$output"
    ~/src/keybase/node-forge-sigchain/bin/main.js --team --format iced --pretty < "$input" > "../$output"
done
