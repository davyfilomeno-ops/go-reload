package main

func apostrophe(mots []string) []string {
	dansApostrophe := false

	for i := 0; i < len(mots); i++ {
		if mots[i] == "'" {
			if !dansApostrophe {
				dansApostrophe = true

				if i+1 < len(mots) {
					mots[i] += mots[i+1]
					mots = append(mots[:i+1], mots[i+2:]...)
				}
			} else {
				dansApostrophe = false

				if i > 0 {
					mots[i-1] += mots[i]
					mots = append(mots[:i], mots[i+1:]...)
					i--
				}
			}
		}
	}

	return mots
}
