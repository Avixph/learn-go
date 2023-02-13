package main

func main() {
	// en := englishBot{}
	// sp := spanishBot{}
	en := roBot{language: "english"}
	sp := roBot{language: "spanish"}

	// printGreeting(en)
	// printGreeting(sp)
	printGreeting(en)
	printGreeting(sp)
}
