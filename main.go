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

var fn = map[rune]string{}

func error(format string, a ...interface{}) {
	os.Exit(1)
}

func expect(r rune) {
	if prg[cnt] != r {
		error("%c expected: %s", prg[cnt])
	}
	cnt++
}

func readUtil(r rune, name rune, target map[rune]string) {
	var buf = ""
	for cnt < len(prg) && prg[cnt] != r {
		buf = buf + string(prg[cnt])
		cnt++
	}
	target[name] = buf
	cnt++
}

func skip() {
	for cnt < len(prg) && unicode.IsSpace(prg[cnt]) {
		cnt++
	}
}

func evalString(code string, args map[rune]int) int {
	orig := prg
	prg = []rune(code)
	origCnt := cnt
	cnt = 0
	val := eval(args)
	prg = orig
	cnt = origCnt
	return val
}

func eval(args map[rune]int) int {
	skip()

	// Function paramter
	if cnt < len(prg) && unicode.IsLower(prg[cnt]) {
		cnt++
		return args[prg[cnt]]
	}

	// Function definition
	if cnt < len(prg) && unicode.IsUpper(prg[cnt]) && prg[cnt+1] == '[' {
		name := prg[cnt]
		cnt += 2
		readUtil(']', name, fn)
		return eval(args)
	}

	// Function application
	if cnt < len(prg) && unicode.IsUpper(prg[cnt]) && prg[cnt+1] == '(' {
		name := prg[cnt]
		cnt += 2
		newarg := eval(args)
		expect(')')
		return evalString(fn[name], newargs)
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
		x := eval(args)
		y := eval(args)
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
	//prg = []rune("F[+ . .] F(1)")
	for cnt < len(prg) {
		fmt.Printf("%d\n", eval(0))
	}
}
