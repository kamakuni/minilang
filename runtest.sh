#!/bin/bash
go build -o lang main.go

runtest() {
    output=$(./lang "$1")
    if [ "$output" != "$2" ]; then
        echo "$1: $2 excepted, but got $output"
        exit 1
    fi
    echo "$1 => $output"
}

runtest 0 0
runtest 1 1
runtest 99 99
runtest '1 2 3' '1 2 3'

echo OK