package main

import "fmt"

func main() {
	arrayNumbers := []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, n := range arrayNumbers {
		if n % 2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}