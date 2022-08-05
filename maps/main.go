package main

import "fmt"

func main() {
	colors := map [string]string{
		"red": "#ff0000",
		"black": "#000000",
		"white": "#ffffff",
	}

	// var colors map[string]string

	// colors := make(map[int]string)
	// colors[10] = "#ffffff"
	// delete(colors, 10)

	printMap(colors)

	fmt.Println(colors);
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Color:" ,color,"Value:", hex)
	}
}
