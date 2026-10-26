package main

import "strings"

func cap(mot string) string {
	mot = strings.ToLower(mot)
	lettres := []rune(mot)

	if len(lettres) == 0 {
		return ""
	}

	premiereLettre := strings.ToUpper(string(lettres[0]))
	resteDuMot := string(lettres[1:])

	return premiereLettre + resteDuMot
}
