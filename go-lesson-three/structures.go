package main

import "fmt"

type Cat struct {
	name, breed string // breed - порода
	years       int8
	children    []Cat
}

// Методы - это функции, которые работают только с определенными типами!
func (c Cat) makeNoise() { // По правилам Golang 'c' эта переменная названа по первой букве типа Сat
	fmt.Printf("Meow-meow from %+v!\n", c)
}

type Movie struct {
	name string
	year int
}

// type Movie struct {
// running TimeInMinutes uint16
// budget InUsd
//uint64
// }

func (m *Movie) print() {
	fmt.Printf("Фильм \"%s\" %d года выпуска\n", m.name, m.year)
}
func (m *Movie) startMovie() {
	fmt.Printf("Фильм \"%s\" НАЧАЛСЯ! \n", m.name)
}

type Character struct {
	Movie        // Встроенное или анонимное
	name  string // обычное или именованное
	// anotherMovie Movie
}

func (c *Character) print() {
	fmt.Printf("Персонаж \"%s\" из фильма %s (%d.)\n", c.name, c.Movie.name, c.year)
}

type Motobike struct {
}

func (m *Motobike) makeNoise() {
	fmt.Printf("Vroom-vroom from %+v! \n", m)
}
