package main

import (
	"fmt"
	"strings"
)

// Execute: go run main.go < caesar.in

func main() {
	var length, delta int
	var input string

	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	fmt.Printf("length: %d\n", length)
	fmt.Printf("input: %s\n", input)
	fmt.Printf("delta: %d\n", delta)

	solutionOne(input, delta)

	var ret []rune
	for _, ch := range input {
		ret = append(ret, solutionTwo(ch, delta))
	}
	fmt.Println(string(ret))
}

func solutionTwo(r rune, delta int) rune {
	switch {
	case r >= 'A' && r <= 'Z':
			return rotateTwo(r, 'A', delta)
	case r >= 'a' && r <= 'z':
		return rotateTwo(r, 'a', delta)
	default:
		return r
	}
}

func rotateTwo(r rune, base, delta int) rune {
	tmp := int(r) - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}

func solutionOne(input string, delta int)  {
	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ret := ""
	for _, ch := range input {
		switch {
		case strings.IndexRune(alphabetLower, ch) >= 0:
			ret = ret + string(rotateOne(ch, delta, []rune(alphabetLower)))
		case strings.IndexRune(alphabetUpper, ch) >= 0:
			ret = ret + string(rotateOne(ch, delta, []rune(alphabetUpper)))
		default:
			ret = ret + string(ch)
		}
	}
	fmt.Println(ret)
}

func rotateOne(s rune, delta int, key []rune) rune {
	idx := strings.IndexRune(string(key), s)
	idx = (idx + delta) % len(key)
	return key[idx]
}