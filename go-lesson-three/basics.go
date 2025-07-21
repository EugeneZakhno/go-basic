package main

import (
	"fmt"
	importedFood "go-basics/go-lesson-three/food"
)

func main() {
	fmt.Println("Lesson_3: Pointers, Maps, Structures, Methods, Public & Private:")
	runLessons()
}

func runLesson(fn func()) {
	fn()
}

func runLessons() {
	runLesson(runPointer)
	runLesson(runMap)
	runLesson(runStructe)
	runLesson(runPublicPrivate)
}

// Pointers
func printArrayAndPointer(array [3]string, pointer *[3]string) {
	fmt.Println("Массив =", array, "; Указатель на array", pointer, "; Значения указателя pointer =", *pointer)
}

func changeNames(array [3]string, pointer *[3]string) {
	array[0] = "Carrie_White"
	pointer[2] = "Donnie_Darko"

	printArrayAndPointer(array, pointer)
}

func runPointer() {
	var mySalary int64 = 1250
	var yourSalary int64 = mySalary
	fmt.Printf("Моя зарплата: %d, твоя зарплата (копия моей): %d \n", mySalary, yourSalary)

	mySalary += 200
	fmt.Printf("Моя зарплата: %d, твоя зарплата: %d \n", mySalary, yourSalary)

	yourSalary += 100
	fmt.Printf("Моя зарплата: %d, твоя зарплата: %d \n", mySalary, yourSalary)

	var mySavings int64 = 10_000
	var yourSavings *int64 = &mySavings
	fmt.Println("Мои сбережения:", mySavings)
	fmt.Println("Твои сбережения: ", yourSavings)
	fmt.Println("Указатель на мои сбережения", &mySavings)
	fmt.Println("На что ссылаются твои сбережения ", *yourSavings)
	fmt.Println("Сравнение по ссылке: ", &mySavings == yourSavings)
	fmt.Println("Cравнение по значению: ", mySavings == *yourSavings)

	//fmt.Println(mySavings == yourSavings)
	//fmt.Println(&mySavings == *yourSavings)

	mySavings += mySalary
	fmt.Printf("Мои сбережения: %d, твои сбережения: %d \n", mySavings, yourSavings)

	var defaultInt64 int64
	var defaultInt64Pointer *int64
	fmt.Println("defaultInt64 =", defaultInt64, "; defaultInt64Pointer =", defaultInt64Pointer)
	fmt.Println("defaultInt64Pointer is nil?", defaultInt64Pointer == nil)
	// fmt.Println(*defaultInt64Pointer)

	var namesArray [3]string = [3]string{"Dana", "Bob", "Alice"}
	var namesPointer *[3]string = &namesArray
	printArrayAndPointer(namesArray, namesPointer)

	changeNames(namesArray, namesPointer)
	printArrayAndPointer(namesArray, namesPointer)

	names := []string{"Dana", "Alice", "Bob"}
	// var pointer **[] string = &&name
	// var pointer **[] string = &(&names)
	pointer1 := &names
	pointer2 := &pointer1
	fmt.Println("Массив:", names)
	fmt.Println("pointer1", pointer1, "; *pointer1", *pointer1)
	fmt.Println("pointer2", pointer2, "; *pointer2", *pointer2, "; **pointer2 =", **pointer2)

	var matreshka1 string = "MАТРЁШКА"
	var matreshka2 *string = &matreshka1
	var matreshka3 **string = &matreshka2
	fmt.Println("Основная матрёшка:", matreshka1, "; Указатель на неё:", matreshka2, "; Указатель на указатель:", matreshka3)
	fmt.Println("Oсновная матрёшка:", matreshka1, "; значение указателя на неё:", *matreshka2, "; Значение указателя на указатель:", **matreshka3)

	**matreshka3 = "НОВАЯ_МАТРЁШКА"
	fmt.Println("Oсновная матрёшка:", matreshka1, "; указатель на неё:", matreshka2, "; Указатель на указатель:", matreshka3)
	fmt.Println("Основная матрёшка:", matreshka1, "; Значение указателя на неё:", *matreshka2, "; Значение указателя на указатель:", **matreshka3)

	var matreshka4 ***string = &matreshka3
	var matreshka5 ****string = &matreshka4
	fmt.Println(matreshka4, ***matreshka4, matreshka5, ****matreshka5)
}

// Map
func printBoolStringMap(mapBoolString map[bool]string) {
	fmt.Println("\n mapBoolString", mapBoolString, ", len", len(mapBoolString), "\n mapBoolString[true] = [",
		mapBoolString[true], "], mapBoolString[false] [", mapBoolString[false], "]")
}

func printStringFloat64Map(mapStringFloat64 map[string]float64) {
	fmt.Println("\n mapBoolString", mapStringFloat64, ", len", len(mapStringFloat64))
}

func runMap() {
	var map1 map[string]float64
	fmt.Println("map1 =", map1, "; len (map1) =", len(map1))
	map1["key1"] = 10
	map2 := make(map[string]float64) // без capacity
	fmt.Println("map2 =", map2, "; len (map2) =", len(map2))

	map2["key1"], map2["key2"] = 100, 200
	fmt.Println(map2["key1"])
	fmt.Println(map2["key2"])
	fmt.Println("map2 =", map2, "; len (map2) =", len(map2))

	map3 := make(map[string]float64, 5) // c capacity
	fmt.Println("map3 =", map3, "; len (map3) =", len(map3))
	// fmt.Println(cap(mapStringFloat64))

	map3["key1"], map3["key2"] = 1000, 2000
	fmt.Println(map3["key1"])
	fmt.Println("map3", map3, "; len (map3) =", len(map3))

	mapBoolString := make(map[bool]string, 1)
	printBoolStringMap(mapBoolString)

	mapBoolString[true] = "Tрава зелёная"
	printBoolStringMap(mapBoolString)

	mapBoolString[true] = "Я подписался на канал"
	printBoolStringMap(mapBoolString)

	mapBoolString[false] = "Земля плоская"
	printBoolStringMap(mapBoolString)

	prices := map[string]float64{
		"Вино":    20.90,
		"Шоколад": 2.10,
		"Платье":  150.00,
		//"Вино":    45.99,
	}

	printStringFloat64Map(prices)

	prices["Вино"] = 45.99
	prices["Пельмени"] = 9.80
	printStringFloat64Map(prices)

	delete(prices, "Вино")
	printStringFloat64Map(prices)

	dressPrice, ok := prices["Платье"]
	fmt.Println("Платье найдено в карте:", ok)
	fmt.Println("Цена платья:", dressPrice)

	vinePrice, ok := prices["Вино"]
	fmt.Println("Вино найдено в карте:", ok)
	fmt.Println("Цена вина:", vinePrice)

	for good, price := range prices {
		fmt.Println("Товар", good, "по цене", price)
	}
	for _, price := range prices { // _ это skip
		fmt.Println("Цена:", price)
	}
	for good := range prices {
		fmt.Println("Товар:", good)
	}
}

// Structures
func changeCharactersNames(character Character, pointer *Character) {
	character.name = "Vincent Vega"
	pointer.name = "Mia Wallace"
}

func runStructe() {
	var emptyCat1 Cat
	emptyCat2 := Cat{}
	fmt.Println("emptyCat1: ", emptyCat1, "emptyCat2 ", emptyCat2)

	var emptyCatPointer1 *Cat = new(Cat)
	var emptyCatPointer2 *Cat = &Cat{}
	fmt.Println("emptyCatPointer1: ", emptyCatPointer1, ", emptyCatPointer2: ", emptyCatPointer2)

	breedSphinx := "Cфинкс"
	breedSiams := "Сиамский"
	var catBosch Cat = Cat{
		name:     "Босх",
		breed:    breedSphinx,
		years:    3,
		children: []Cat{},
	}
	fmt.Println(catBosch)

	var catVermeer Cat = Cat{"Вермеер", breedSphinx, 5, []Cat{
		{"Магритт", breedSphinx, 1, nil},
		{"Боттичелли", breedSphinx, 2, nil},
	}}
	fmt.Println(catVermeer)

	var vasiliy Cat = Cat{
		name:  "Василий",
		breed: breedSiams,
		years: 6,
		// Owner: "Lora Palmer"
	}

	vasiliy.breed = breedSiams
	fmt.Println(vasiliy)
	fmt.Printf(" Так удобнее указываем ключ значение %+v \n", vasiliy)
	fmt.Println("Порода Василия:", vasiliy.breed)

	// const marfa Cat   ВАЖНО: структуру нельзя присвоить константе, т.е. структура не может быть константой!!!!

	biba, boba := Cat{"Биба", breedSiams, 3, nil},
		Cat{"Боба", breedSiams, 3, nil}

	boba.children = []Cat{{"Боба II", breedSiams, 1, nil}}
	biba.name = "Boba"
	fmt.Printf("Дети Бобы: %+v\n", boba.children)

	vasiliy.children = append(vasiliy.children, biba)
	fmt.Printf("%+v \n", vasiliy)

	vasiliy.children = append(vasiliy.children, boba)
	fmt.Printf("%+v \n", vasiliy)

	boba.name = "Боба І"
	fmt.Printf("%+v \n", boba)
	fmt.Printf("%+v \n", vasiliy)
	vasiliy.children[1].name = "Боба Старший"
	fmt.Printf("%+v \n", vasiliy)
	fmt.Printf("%+v \n", boba)

	vasiliy.makeNoise()

	var moto Motobike = Motobike{}
	fmt.Println("Пустой мотоцикл =", moto)
	moto.makeNoise()

	moto.makeNoise()

	pulpFiction := Movie{"Криминальное чтиво", 1994}
	pulpFiction.print()
	pulpFiction.startMovie()

	fmt.Println(pulpFiction == Movie{"Криминальное чтиво", 1994})
	fmt.Println(pulpFiction == Movie{"Бой с тенью", 2005})

	var emptyMovie1 Movie
	var emptyMovie2 Movie

	fmt.Println(emptyMovie1 == pulpFiction)
	fmt.Println(emptyMovie1 == emptyMovie2)

	// Важно: мы не можем сравнивать структуры и nil, мы можем сравнивать pointers и nil !
	// fmt.Println(nilMovie1 == nil)

	// Важно: мы не можем сравнивать структуры в которых содержаться другие встроенные объекты, как например Cat {}
	// fmt.Println(vasiliy == biba)

	vincentVega := Character{
		Movie: pulpFiction,
		name:  "Винсент Вега",
	}

	miawallace := Character{
		pulpFiction,
		"Мия Уоллес", //Если переношу скобку, то надо запятую указать, а если не переношу, то не надо))) ха-ха, прикол!
	}
	miawallace.print()
	miawallace.startMovie()

	// miawallace := Character {pulpFiction, "Мия Уоллес", emptyMovie1}
	// fmt.Println(miawallace.Movie.name)
	// fmt.Println(miawallace.anotherMovie.name)

	var emptyCharacter Character

	fmt.Println(vincentVega)
	fmt.Println(miawallace)
	fmt.Println(emptyCharacter)

	fmt.Println("Год через фильм:", vincentVega.Movie.year)
	fmt.Println("Год напрямую:", vincentVega.year)

	changeCharactersNames(vincentVega, &miawallace)
	fmt.Println(vincentVega, miawallace)
	newMovie := Movie{"Криминальное чтиво", 1994}
	movieAndRateMap := map[Movie]int8{
		pulpFiction: 5,
		newMovie:    4,
		{"Криминальное чтиво", 1994}: 3}

	fmt.Println(movieAndRateMap)

	moviePointerAndRateMap := map[*Movie]int8{
		&pulpFiction: 5,
		&newMovie:    4,
		{"Криминальное чтиво", 1994}: 3}

	fmt.Println(moviePointerAndRateMap)
}

// Public & Private  // Большая буква это Публичный доступ, маленькая первая буква это приватный доступ.
func runPublicPrivate() {
	// publicFruit: importedFood. Fruit {"Манго", 67}
	publicFruit := importedFood.Fruit{Name: "яблоко"}

	// fmt.Println(publicFruit.calories)
	fmt.Println(publicFruit.Name)

	// publicFruit.bite(3)
	fmt.Println(publicFruit.GetDescription())
	// = importedFood.fastFood {"Шаверма", 800}
}

// Main
func cleanScreen() {
}

func pause() {
}
