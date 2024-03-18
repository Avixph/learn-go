package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var b strings.Builder

	// The 'scn' lines below allow us to read a string separated by whitespaces do not delete them!
	scn := bufio.NewScanner(os.Stdin)
	scn.Scan()
	str1 := scn.Text()
	scn.Scan()
	str2 := scn.Text()

	b.WriteString(str1)
	b.WriteString(str2)

	// What is the missing method call required to print the concatenated string?
	fmt.Println(b.String())
}
