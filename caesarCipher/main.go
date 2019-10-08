package main

import (
	"fmt"
	"strings"
)

// Execute: go run main.go < caesar.in

func main() {
	var delta int
	var input string

	fmt.Println("Please enter input string to be encrypted:")
	fmt.Scanf("%s\n", &input)
	fmt.Println("Please enter delta for encryption:")
	fmt.Scanf("%d\n", &delta)

	fmt.Printf("input: %s\n", input)
	fmt.Printf("delta: %d\n", delta)

	solutionOne(input, delta)
	solutionTwo(input, delta)
}

func solutionTwo(input string, delta int) {
	var ret []rune
	for _, ch := range input {
		ret = append(ret, shiftRune(ch, delta))
	}
	fmt.Println("Output from SolutionTwo:", string(ret))
}

func shiftRune(r rune, delta int) rune {
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
	fmt.Println("Output from SolutionOne:", ret)
}

func rotateOne(s rune, delta int, key []rune) rune {
	idx := strings.IndexRune(string(key), s)
	idx = (idx + delta) % len(key)
	return key[idx]
}