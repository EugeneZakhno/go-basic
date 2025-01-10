package main

import "fmt"

func main() {
	fmt.Println("Lesson_3:")
	runLessons()
	runMap()
}
func runLesson(m func()) {

}

func runLessons() {
	runLesson(runPointer)
	runLesson(runMap)
	//runLesson(runStructure)
	//		runLesson(runPublicPrivate)
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
	fmt.Printf("Моя зарплата: %d, твоя зарплата: %d \n", mySalary, yourSalary)

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
	//fmt.Println(&mySavings == yourSavings)

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
}

func runMap() {
	var map1 map[string]float64
	fmt.Println("map1 =", map1, "; len (map1) =", len(map1))
	// map1 ["key1"] = 10
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

func runStructe() {

}
