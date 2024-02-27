package main

import "fmt"

// Update the parameter list of the countryCapital function below:
func countryCapital(capital, country string) string {
	return capital + " is the capital of " + country
}

// DO NOT modify or delete the contents within the main function!
func main() {
	var capital, country string
	_, err := fmt.Scanln(&capital, &country)
	if err != nil {
		return
	}

	fmt.Println(countryCapital(capital, country))
}
