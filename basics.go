package main //Обязательно main

import (
	"fmt"
)

// Global var
var banan string = "Mango"
var favoriteNum int = 100

func main() { //обязательно main

	var favoriteBook string = "Сияние"
	//Local var
	fmt.Println("Hello Go" + "!")
	fmt.Println(favoriteBook)
	favoriteBook = "}{уета"
	fmt.Println(favoriteBook)
}
