package main

import "fmt"

func versHex(mot string) string {
	resultat := ""

	for i := 0; i < len(mot); i++ {
		resultat += fmt.Sprintf("%02x", mot[i])
	}

	return resultat
}
