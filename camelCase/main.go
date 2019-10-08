package main

import (
	"fmt"
	"strings"
)

/*
	Execute code: go run main.go
	Execute code and take input from camelCase.in: go run main.go < camelCase.in
	Execute code and move output to camelCase.out: go run main.go > camelCase.out
	All the above: go run main.go < camelCase.in > camelCase.out
 */

/*
	BYTE/RUNE explanation

	min, max := 'A', 'Z'

	// Any string is internally a BYTE/RUNE array
	b := []byte(input)
	fmt.Println(b)

	for _, ch := range input {
		fmt.Print(ch, " (index ", i, ")") //Prints same as above
		// Try providing an EMOJI as input along with normal string
		if ch >= min && ch <= max {
			// It is a capital letter !
			answer++
		}
	}
*/

func main()  {
	var input string
	answer := 1
	fmt.Println("Please enter camelCase string:")
	fmt.Scanf("%s\n", &input)
	fmt.Println("Input received:", input)

	for _, ch := range input {
		str := string(ch)

		if strings.ToUpper(str) == str {
			// It is a capital letter !
			answer++
		}
	}

	fmt.Println("Number of words are:", answer)
}