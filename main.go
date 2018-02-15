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

func error(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a)
	os.Exit(1)
}

func skip() {
	for cnt < len(prg) && unicode.IsSpace(prg[cnt]) {
		cnt++
	}
}

func eval() int {
	skip()

	if cnt < len(prg) && unicode.IsDigit(prg[cnt]) {

		var val = int(prg[cnt] - '0')
		cnt++
		for cnt < len(prg) && unicode.IsDigit(prg[cnt]) {
			val = val*10 + int(prg[cnt]-'0')
			cnt++
		}
		return val
	}

	//	if cnt < len(prg) && (prg[cnt] == '+' || prg[cnt] == '-') {
	if cnt < len(prg) && strings.ContainsRune("+-*/", prg[cnt]) {
		op := prg[cnt]
		cnt++
		switch op {
		case '+':
			return eval() + eval()
		case '-':
			return eval() - eval()
		case '*':
			return eval() * eval()
		case '/':
			return eval() / eval()
		}
	}
	error("invalid value %c", prg[cnt])
	return 0
}

func main() {
	prg = []rune(os.Args[1])
	for cnt < len(prg) {
		fmt.Printf("%d\n", eval())
	}
}
