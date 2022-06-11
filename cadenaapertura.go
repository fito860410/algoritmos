package main

import (
	"strings"
)

//función que evalúa si una expresión está equilibrada o no
func equilibrada(valor string) bool {

	var char = strings.Split(valor, "")
	slice1 := []string{}

	for i := 0; i < len(valor); i++ {
		if char[i] == "(" || char[i] == "[" || char[i] == "{" {
			slice1 = append(slice1, char[i])
		} else {
			if char[i] == ")" {
				if len(slice1) == 0 {
					return false
				} else if slice1[len(slice1)-1] != "(" {
					return false
				} else {
					slice1 = slice1[:len(slice1)-1]
				}
			} else {
				if char[i] == "]" {

					if len(slice1) == 0 {
						return false
					} else if slice1[len(slice1)-1] != "[" {
						return false
					} else {
						slice1 = slice1[:len(slice1)-1]
					}
				} else {
					if char[i] == "}" {
						if len(slice1) == 0 {
							return false
						} else if slice1[len(slice1)-1] != "{" {
							return false
						} else {
							slice1 = slice1[:len(slice1)-1]
						}
					}
				}
			}
		}
	}

	if len(slice1) == 0 {
		return true
	} else {
		return false
	}

}

func main() {
	r := equilibrada("(8 + 4) + 2733} ")
	println(r)
}
