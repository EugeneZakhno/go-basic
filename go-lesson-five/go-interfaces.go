package main

import "fmt"

// Ehot

type Racoon struct {
	name string
}

func (r Racoon) GetName() string {
	return r.name
}
func (r Racoon) Print() {
	fmt.Println("Меня зовут", r.GetName())
}
func (r Racoon) InvestigateGarbage() {
	fmt.Println("Енот: Поиски сокровищ в мусорном баке в самом разгаре")
}

// func (r *Racoon) Print() {
// fmt.Println("Меня зовут", r.GetName())
// }

// Кукушка
type Cuckoo struct {
	name string
}

func (c Cuckoo) GetName() string {
	return c.name
}
func (c Cuckoo) Print() {
	fmt.Println("Я кукушка по имени", c.GetName())
}

func (c Cuckoo) GetYearLeft() int {
	// return int(rand.Int31n(60) + 30) // Для интерфейсов
	return len(c.name) * 10 // Для многопоточности
}

// Олень
type Deer struct {
	name string
}

func (d Deer) GetName() string {
	return d.name
}
func (d Deer) Print() {
	fmt.Println("Я олень по имени", d.GetName())
}
func (d Deer) GetPicture() string {
	return string('\U0001F98C')
}

// Животное
type Animal interface {
	GetName() string
	Print()
	//Name string
}
