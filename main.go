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

//var fn []string
var fn map[rune]string = map[rune]string{}

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

func readUtil(r rune, name rune, target map[rune]string) {
	var buf = ""
	fmt.Println("******************************")
	fmt.Println(string(r))
	fmt.Printf("%d", cnt)
	fmt.Println(string(prg[cnt]))
	fmt.Println("******************************")
	for prg[cnt] != r {
		fmt.Println("loooooooooooooop")
		fmt.Println(buf)
		buf = buf + string(prg[cnt])
		cnt++
	}
	fmt.Println(buf)
	target[name] = buf
}

func skip() {
	for cnt < len(prg) && unicode.IsSpace(prg[cnt]) {
		cnt++
	}
}

func evalString(code string, arg int) int {
	orig := prg
	origCnt := cnt
	cnt = 0
	val := eval(arg)
	prg = orig
	cnt = origCnt
	return val
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
		fmt.Println(string(name))
		readUtil(']', name, fn)
		return eval(arg)
	}

	// Function application
	if cnt < len(prg) && unicode.IsUpper(prg[cnt]) && prg[1] == '(' {
		name := prg[cnt]
		cnt += 2
		newarg := eval(arg)
		expect(')')
		return evalString(fn[name-'A'], newarg)
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
