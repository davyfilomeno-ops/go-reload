package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	texte, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Erreur lecture :", err)
		return
	}

	mots := strings.Fields(string(texte))
	for i := 0; i < len(mots); i++ {
		nom := ""
		nombreMots := 1
		nombreTokens := 1

		switch mots[i] {
		case "(up)":
			nom = "up"
		case "(low)":
			nom = "low"
		case "(cap)":
			nom = "cap"
		case "(up,", "(low,", "(cap,":
			nom = strings.Trim(mots[i], "(,")
			nombreMots, _ = strconv.Atoi(strings.TrimSuffix(mots[i+1], ")"))
			nombreTokens = 2
			if i+2 < len(mots) && mots[i+2] == ")" {
				nombreTokens++
			}
		}

		if nom == "" {
			continue
		}

		for j := i - nombreMots; j < i; j++ {
			switch nom {
			case "up":
				mots[j] = up(mots[j])
			case "low":
				mots[j] = low(mots[j])
			case "cap":
				mots[j] = cap(mots[j])
			}
		}

		mots = append(mots[:i], mots[i+nombreTokens:]...)
		i--
	}

	resultat := strings.Join(mots, " ")
	fmt.Println(resultat)

	err = os.WriteFile("output.txt", []byte(resultat), 0644)
	if err != nil {
		fmt.Println("Erreur écriture :", err)
		return
	}
}
