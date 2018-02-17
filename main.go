package main

/*
 * https://www.youtube.com/watch?v=JAtN0TGrNE4
 */

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

var prg []rune
var cnt int
var fn []string

func error(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a)
	os.Exit(1)
}

func expect(r rune) {
	if prg[cnt] != r {
		error("%c expected: %s", prg[cnt])
	}
	cnt++
}

func eval_string() {

}

func readUtil(r rune, target *string) {
	for prg[cnt] == r {
		*target = *target + string(prg[cnt])
		cnt++
	}
	fmt.Println(*target)
}

func skip() {
	for cnt < len(prg) && unicode.IsSpace(prg[cnt]) {
		cnt++
	}
}

func eval(arg int) int {
	skip()

	// Function paramter
	if cnt < len(prg) && prg[cnt] == '.' {
		cnt++
		return arg
	}

	// Function definition
	if cnt < len(prg) && unicode.IsUpper(prg[cnt]) && prg[1] == '[' {
		name := prg[cnt]
		cnt += 2
		readUtil(']', &fn[name-'A'])
		return eval(arg)
	}

	// Function application
	if cnt < len(prg) && unicode.IsUpper(prg[cnt]) && prg[1] == '(' {
		name := cnt
		p += 2

		skip()

	}

	// Literal numbers
	if cnt < len(prg) && unicode.IsDigit(prg[cnt]) {

		var val = int(prg[cnt] - '0')
		cnt++
		for cnt < len(prg) && unicode.IsDigit(prg[cnt]) {
			val = val*10 + int(prg[cnt]-'0')
			cnt++
		}
		return val
	}

	// Aritmetic operators
	if cnt < len(prg) && strings.ContainsRune("+-*/", prg[cnt]) {
		op := prg[cnt]
		cnt++
		x := eval(arg)
		y := eval(arg)
		switch op {
		case '+':
			return x + y
		case '-':
			return x - y
		case '*':
			return x * y
		case '/':
			return x / y
		}
	}
	error("invalid value %c", prg[cnt])
	return 0
}

func main() {
	prg = []rune(os.Args[1])
	for cnt < len(prg) {
		fmt.Printf("%d\n", eval(0))
	}
}
