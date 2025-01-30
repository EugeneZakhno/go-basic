package food

import "fmt"

func demonstrate() {
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
