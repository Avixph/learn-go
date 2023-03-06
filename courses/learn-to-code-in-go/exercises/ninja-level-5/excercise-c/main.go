package main

import "fmt"

func main() {
	fordF150 := truck{vehicle: vehicle{doors: 2, color: "black"}, fourWheel: true}
	audiA4 := sedan{vehicle: vehicle{doors: 4, color: "red"}, luxury: true}

	fmt.Println(fordF150)
	fmt.Println(fordF150.color)
	fmt.Println(audiA4)
	fmt.Println(audiA4.color)
}
