package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	texte, cheminEntree, err := lireEntree()
	if err != nil {
		fmt.Println("Erreur lecture :", err)
		return
	}

	mots := strings.Fields(string(texte))
	mots = apostrophe(mots)
	mots, err = appliquerCommandes(mots)
	if err != nil {
		fmt.Println("Erreur traitement :", err)
		return
	}

	resultat := strings.Join(mots, " ")
	resultat = corrigerPonctuation(resultat)
	fmt.Println(resultat)

	cheminSortie := filepath.Join(filepath.Dir(cheminEntree), "output.txt")
	err = os.WriteFile(cheminSortie, []byte(resultat), 0644)
	if err != nil {
		fmt.Println("Erreur écriture :", err)
	}
}

func lireEntree() ([]byte, string, error) {
	chemins := []string{"input.txt", filepath.Join("go-reload", "input.txt")}

	for _, chemin := range chemins {
		texte, err := os.ReadFile(chemin)
		if err == nil {
			return texte, chemin, nil
		}
		if !os.IsNotExist(err) {
			return nil, "", err
		}
	}

	return nil, "", fmt.Errorf("input.txt est introuvable")
}

func appliquerCommandes(mots []string) ([]string, error) {
	mots = separerPonctuationCommandes(mots)

	for indexCommande := 0; indexCommande < len(mots); indexCommande++ {
		nom, nombreDeMots, nombreDeTokens, commandeReconnue, err := lireCommande(mots, indexCommande)
		if err != nil {
			return nil, err
		}
		if !commandeReconnue {
			continue
		}

		if nombreDeMots < 1 || nombreDeMots > indexCommande {
			return nil, fmt.Errorf(
				"la commande %q demande %d mot(s), mais seuls %d sont disponibles",
				mots[indexCommande], nombreDeMots, indexCommande,
			)
		}

		for indexMot := indexCommande - nombreDeMots; indexMot < indexCommande; indexMot++ {
			switch nom {
			case "up":
				mots[indexMot] = up(mots[indexMot])
			case "low":
				mots[indexMot] = low(mots[indexMot])
			case "cap":
				mots[indexMot] = cap(mots[indexMot])
			}
		}

		mots = append(mots[:indexCommande], mots[indexCommande+nombreDeTokens:]...)
		indexCommande--
	}

	for indexCommande := 0; indexCommande < len(mots); indexCommande++ {
		if mots[indexCommande] != "(hex)" && mots[indexCommande] != "(bin)" {
			continue
		}
		if indexCommande == 0 {
			return nil, fmt.Errorf("la commande %q n'a aucun mot à convertir", mots[indexCommande])
		}

		if mots[indexCommande] == "(hex)" {
			mots[indexCommande-1] = versHex(mots[indexCommande-1])
		} else {
			mots[indexCommande-1] = versBin(mots[indexCommande-1])
		}

		mots = append(mots[:indexCommande], mots[indexCommande+1:]...)
		indexCommande--
	}

	return mots, nil
}

func separerPonctuationCommandes(mots []string) []string {
	var resultat []string

	for _, mot := range mots {
		commande := strings.TrimRight(mot, ".,!?:;")
		ponctuation := strings.TrimPrefix(mot, commande)

		if ponctuation != "" && estCommande(commande) {
			resultat = append(resultat, commande, ponctuation)
		} else {
			resultat = append(resultat, mot)
		}
	}

	return resultat
}

func estCommande(mot string) bool {
	switch mot {
	case "(up)", "(low)", "(cap)", "(hex)", "(bin)":
		return true
	}

	for _, nomCommande := range []string{"up", "low", "cap"} {
		if strings.HasPrefix(mot, "("+nomCommande+",") && strings.HasSuffix(mot, ")") {
			return true
		}
	}

	return false
}

func lireCommande(mots []string, index int) (string, int, int, bool, error) {
	mot := mots[index]

	switch mot {
	case "(up)":
		return "up", 1, 1, true, nil
	case "(low)":
		return "low", 1, 1, true, nil
	case "(cap)":
		return "cap", 1, 1, true, nil
	}

	nomCommande := ""
	for _, nomPossible := range []string{"up", "low", "cap"} {
		if strings.HasPrefix(mot, "("+nomPossible+",") {
			nomCommande = nomPossible
			break
		}
	}
	if nomCommande == "" {
		return "", 0, 0, false, nil
	}

	nombreTexte := strings.TrimPrefix(mot, "("+nomCommande+",")
	nombreDeTokens := 1

	if nombreTexte == "" {
		if index+1 >= len(mots) {
			return "", 0, 0, true, fmt.Errorf("nombre de mots manquant dans %q", mot)
		}
		nombreTexte = mots[index+1]
		nombreDeTokens++
	}

	if strings.HasSuffix(nombreTexte, ")") {
		nombreTexte = strings.TrimSuffix(nombreTexte, ")")
	} else if index+nombreDeTokens < len(mots) && mots[index+nombreDeTokens] == ")" {
		nombreDeTokens++
	} else {
		return "", 0, 0, true, fmt.Errorf("commande mal formée : %q", mot)
	}

	nombreDeMots, err := strconv.Atoi(nombreTexte)
	if err != nil {
		return "", 0, 0, true, fmt.Errorf("nombre invalide dans %q : %w", mot, err)
	}

	return nomCommande, nombreDeMots, nombreDeTokens, true, nil
}
