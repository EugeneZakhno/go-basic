package food

import "fmt"

type Fruit struct { //public type
	Name     string
	calories int16
}

func (f *Fruit) GetDescription() string { // public метод, т.к. с большой буквы начинается название
	return fmt.Sprintf("Я фрукт \"%s\", во мне % ккал\n", f.Name, f.calories)
}

func (f *Fruit) bite(times int8) { // private метод, т.к. с маленькой буквы начинается название
	for i := 1; i <= int(times); i++ {
		fmt.Printf("Ты кусаешь фрукт \"%s\"", f.Name)
	}
}

type fastFood struct { // private type
	Name     string
	calories int16
}

func (f *fastFood) GetDescription() string {
	return fmt.Sprintf("Я фаст-фуд \"%s\", во мне %d ккал\n", f.Name, f.calories)
}

func (f *fastFood) bite(times int8) {
}

func main() {
	publicFruit := Fruit{"Манго", 67}
	fmt.Println(publicFruit.Name, publicFruit.calories)
	fmt.Println(publicFruit.GetDescription())
	publicFruit.bite(3)
	privateFastFood := fastFood{"Шаверма", 800}

	fmt.Println(privateFastFood.Name, privateFastFood.calories)
	fmt.Println(privateFastFood.Name, privateFastFood.calories)
	fmt.Println(privateFastFood.GetDescription())
	privateFastFood.bite(5)
}
