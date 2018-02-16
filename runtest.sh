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

echo '=== basic ==='
runtest 0 0
runtest 1 1
runtest 99 99
runtest '1 2 3' '1
2
3'

echo '=== arithmetic operators ==='
runtest '+ 1 2' 3
runtest '+ 100 5' 105
runtest '- 5 1' 4
runtest '- 1 4' -3
runtest '* 3 5' 15
runtest '/ 20 5' 4
runtest '+ + + 1 2 3 4' 10
runtest '+ 1 + 2 + 3 4' 10
runtest '+ 2 * 4 3' 14

echo OK
