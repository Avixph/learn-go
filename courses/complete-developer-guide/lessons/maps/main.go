package main

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)
	colors := map[string]string{
		"red":    "#ff0000",
		"yellow": "#ffff00",
		"blue":   "#0000ff",
	}
	colors["orange"] = "#ffa500"
	colors["green"] = "#00ff00"
	colors["violet"] = "#8f00ff"
	colors["white"] = "#ffffff"
	delete(colors, "white")
	printMap(colors)
}
