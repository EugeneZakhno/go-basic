package main

import "fmt"

type Good struct {
	Id   int64
	Name string
}
type Employee struct {
	Id             int64
	Name, Position string
}

type Document struct {
	Id     int64
	Type   string
	Goods  []Good
	Author *Employee
}

func (d *Document) Print() {
	fmt.Println("----------------------------------")
	fmt.Printf("| Документ \"%s\" %d\n", d.Type, d.Id)
	fmt.Println(" ~ ~ ~ ~ ~ ~ ~~~~~~~~~~")
	for index, good := range d.Goods {
		fmt.Printf("| %d) %5 (код %d)\n", index+1, good.Name, good.Id)
	}
	fmt.Println(" ~ ~ ~ ~ ~ ~ ~~~~~~~~~~")
	fmt.Printf("| Автор: %s %s\n", d.Author.Position, d.Author.Name)
	fmt.Println("------------------")
	fmt.Println()
}
