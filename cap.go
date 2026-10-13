package main

import "strings"

func cap(mot string) string {
	return strings.ToUpper(mot[:1]) + strings.ToLower(mot[1:])
}	