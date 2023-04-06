package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	vowels := []string{"a", "e", "i", "o", "u", "y"}
	consonants := []string{"b", "c", "d", "f", "g", "j", "k", "l", "m", "n", "p", "q", "s", "t", "v", "x", "z"}
	honorifics := []string{"mr", "mrs", "miss", "ms", "mx", "sir", "dame", "dr", "cllr", "lady", "lord"}
	name := ""

	for _, honorific := range honorifics {
		for _, consonant := range consonants {
			for _, vowel := range vowels {
				name = fmt.Sprintf("%v. %v%v%v%v", honorific, consonant, vowel, vowel, consonant)
			}
		}
	}
	fmt.Println("hello", name)
	wg.Done()
}

func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println(i + 1)
	}
	wg.Done()
}
