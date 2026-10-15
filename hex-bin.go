package main 

import (
	"strconv"
)	

func conversion(mots []string) []string {

    for i := 0; i < len(mots); i++ {

        if i > 0 && (mots[i] == "(hex)" || mots[i] == "(bin)") {

            base := 16

            if mots[i] == "(bin)" {
                base = 2
            }

            nombre, err := strconv.ParseInt(mots[i-1], base, 64)

            if err == nil {
                mots[i-1] = strconv.FormatInt(nombre, 10)
            }

            mots = append(mots[:i], mots[i+1:]...)
            i--
        }
    }

    return mots
}