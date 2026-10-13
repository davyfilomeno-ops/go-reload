package main 

func estPonctuation(caractere rune) bool {
	return caractere == '.' ||
		caractere == ',' ||
		caractere == '!' ||
		caractere == '?' ||
		caractere == ':' ||
		caractere == ';'
}

func corrigerPonctuation(texte string) string {

	resultat := ""

	for i := 0; i < len(texte); i++ {

		if texte[i] == ' ' && i+1 < len(texte) && estPonctuation(rune(texte[i+1])) {
			continue
		}

		resultat += string(texte[i])
	}

	return resultat
}