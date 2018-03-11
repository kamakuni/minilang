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
var fn = map[rune]string{}

func error(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a)
	os.Exit(1)
}

func expect(r rune) {
	fmt.Printf("rune:%d", r)
	if prg[cnt] != r {
		fmt.Printf("rune:%d", prg[cnt])
		error("%c expected: %s", prg[cnt])
	}
	cnt++
}

func readUtil(r rune, name rune, target map[rune]string) {
	var buf = ""
	fmt.Println("*******************************")
	for cnt < len(prg) && prg[cnt] != r {
		fmt.Printf("prg:%s\n", string(prg[cnt]))
		fmt.Printf("r:%s\n", string(r))
		buf = buf + string(prg[cnt])
		cnt++
		fmt.Printf("cnt:%d\n", cnt)
	}
	fmt.Println("*******************************")
	fmt.Printf("function name:%s\n", string(name))
	fmt.Printf("function:%s\n", buf)
	target[name] = buf
	cnt++
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
		fmt.Println("Function param")
		cnt++
		return arg
	}

	// Function definition
	if cnt < len(prg) && unicode.IsUpper(prg[cnt]) && prg[1] == '[' {
		name := prg[cnt]
		cnt += 2
		readUtil(']', name, fn)
		return eval(arg)
	}
	fmt.Println("---------------------")
	fmt.Printf("prg:%s\n", string(prg[cnt]))
	fmt.Printf("cnt:%d\n", string(prg[cnt+1]))
	fmt.Println("---------------------")
	// Function application
	if cnt < len(prg) && unicode.IsUpper(prg[cnt]) && prg[cnt+1] == '(' {
		fmt.Printf("func apply:%d\n", cnt)
		name := prg[cnt]
		fmt.Printf("name:%d\n", name)
		cnt += 2
		fmt.Printf("cnt3:%d\n", cnt)
		newarg := eval(arg)
		expect(')')
		return evalString(fn[name], newarg)
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
