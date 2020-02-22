package main

import (
	"bytes"
	"fmt"
	"strings"
)

func ToCamelCase(s string) string {
	var str bytes.Buffer
	var do bool

	for _, ch := range s {
		if ch != '_' && ch != '-' {
			if do {
				ch = rune(strings.ToUpper(string(ch))[0])
			}

			str.WriteRune(ch)
			do = false
		} else {
			do = true
		}
	}

	return str.String()
}

func main() {
	fmt.Println(ToCamelCase("the-stealth-warrior"))
	fmt.Println(ToCamelCase("The_Stealth_Warrior"))
}
