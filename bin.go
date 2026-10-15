package main

import "fmt"

func versBin(mot string) string {
	resultat := ""

	for i := 0; i < len(mot); i++ {
		resultat += fmt.Sprintf("%08b", mot[i])
	}

	return resultat
}
