package main

func estPonctuation(caractere rune) bool {
	switch caractere {
	case '.', ',', '!', '?', ':', ';':
		return true
	default:
		return false
	}
}

func corrigerPonctuation(texte string) string {
	resultat := ""

	for i := 0; i < len(texte); i++ {
		caractere := rune(texte[i])

		if caractere == ' ' && i+1 < len(texte) {
			prochainCaractere := rune(texte[i+1])
			if estPonctuation(prochainCaractere) {
				continue
			}
		}

		resultat += string(caractere)
	}

	return resultat
}
