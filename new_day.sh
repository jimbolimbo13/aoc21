#!/usr/bin/env bash
set -euo pipefail

if [ "$#" -eq 0 ] ; then
    echo "What day is it?"
    read today
else
    today=$1
fi

cp utils/default.go "utils/day$today.go"
touch "data/day"$today"_example.txt"
touch "data/day"$today"_input.txt"

sed -i "s/XX/$today/g" utils/day$today.go
