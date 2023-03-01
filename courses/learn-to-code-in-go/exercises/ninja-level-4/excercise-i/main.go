package main

import "fmt"

func main() {
	jb := []string{"Shaken, not stirred", "Martinis", "Woman"}
	mm := []string{"James Bond", "Literature", "Computer Science"}
	dn := []string{"Being evil", "Ice cream", "unsets"}
	jbCharacters := map[string][]string{
		"bond_james":      jb,
		"moneypenny_miss": mm,
		"no_dr":           dn,
	}

	jbCharacters["fleming_ian"] = []string{"steak", "cigars", "espionage"}

	for i, characters := range jbCharacters {
		fmt.Println("character:", i)
		for j, character := range characters {
			fmt.Println(j+1, character)
		}
	}
}
