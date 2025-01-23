package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Lesson_04_ : https://www.youtube.com/watch?v=Gk6BXGF7mko&list=PLcLNZ0ymDkp7XKi7Yg-4maj4kXBV-mZpF&index=4")
	runLessons()
}

const (
	TEXT_FILE_NAME = "go-lesson-four/documents/price-list.txt" // не забывай прописывать путь от корневой папки (начальной, т.е. название проекта (репозитория))
	JSON_FILE_NAME = "go-lesson-four/documents/documents.json"
)

type Author struct {
}

func runLessons() {
	runLesson(runMath)
	//runLesson(runRandom)
	//runLesson(runSlicesSort)
	//runLesson(runDefer)
	//runLesson(runErrors)
	runLesson(runReadWriteToFile)
	//runLesson(runReadFromConsole)
}

func runLesson(m func()) {
	m() // Вызов функции m
}

// Math
func demonstrateMaxMinForTypes() {
	var maxInt8Value int8 = math.MaxInt8
	var maxInt16Value int16 = math.MaxInt16
	var maxInt32Value int32 = math.MaxInt32
	var maxInt64Value int64 = math.MaxInt64
	fmt.Println("Max values for int ->  ", maxInt8Value, maxInt16Value, maxInt32Value, maxInt64Value)

	var Int8Value int8
	var Int16Value int16
	var Int32Value int32
	var Int64Value int64
	fmt.Println("Values for int by default ->  ", Int8Value, Int16Value, Int32Value, Int64Value)

	var minInt8Value int8 = math.MinInt8
	var minInt16Value int16 = math.MinInt16
	var minInt32Value int32 = math.MinInt32
	var minInt64Value int64 = math.MinInt64
	fmt.Println("Min values for int ->  ", minInt8Value, minInt16Value, minInt32Value, minInt64Value)

	var minsInt8Value int8 = math.MinInt8
	var minsInt16Value int32 = math.MinInt16   // преобразуется к int32 автоматически под капотом
	var minsInt32Value float32 = math.MinInt32 // преобразуется к float32 автоматически под капотом
	var minsInt64Value int = math.MinInt64
	var minsIntValue float64 = math.MinInt
	fmt.Println(minsInt8Value, minsInt16Value, minsInt32Value, minsInt64Value, minsIntValue)

	//var maxInt16InInt8 int8 = math.MaxInt16 // maxInt16 = 32767, maxInt8 = 127 не может преобразовать, т.к. меньше  (не лезе)
	//var minInt16InInt8 int8 = math.MinInt16 // minInt16 = -32768, minInt8 = -128 не может преобразовать, т.к. меньше (не лезе)
	var maxUint8 uint8 = math.MaxUint8
	var maxUint16 uint16 = math.MaxUint16
	var maxUint32 uint32 = math.MaxUint32
	var maxUint64 uint64 = math.MaxUint64
	var maxUint uint = math.MaxUint
	fmt.Println(maxUint8, maxUint16, maxUint32, maxUint64, maxUint)

	// maxUintInDefaultIn:= math.MaxUint
	// maxUint64InDefaultIn:= math.MaxUint64

	var maxFloat32 float32 = math.MaxFloat32
	var maxFloat64 float64 = math.MaxFloat64

	// var minFloat32 float32 math. MinFloat32
	// var minFloat64 float64 math.MinFloat64

	fmt.Println(maxFloat32, maxFloat64)
}

// Показать минимальные и максимальные значения между типами
func demonstrateMaxMinBetweenValues() {
	var maxFloat float64 = math.Max(1.1, 1.2)
	fmt.Println("max =", maxFloat)

	var minFloat float64 = math.Min(1.1, 1.2)
	fmt.Println("min =", minFloat)

	minValue := math.Min(1, 2)
	fmt.Println("min =", minValue)

	var intValue1, intValue2 int = 1, 2
	// _ = math.Max(intValue1, intValue2) Невозможно передать в функцию Max значения с типом int Golang всеравно преобразует в float64
	_ = math.Max(float64(intValue1), float64(intValue2))

	var float32Value1, float32Value2 float32 = 1.1, 2.2
	// _ = math.Min(float32Value1, float32Value2) Невозможно передать в функцию Max значения с типом float32 a Golang всеравно преобразует в float64
	_ = math.Min(float64(float32Value1), float64(float32Value2))
}

func demonstrateRounding() {
	// Ceil всегда ↑
	var ceil1, ceil2, ceil3 float64 = math.Ceil(2.49), math.Ceil(2.5), math.Ceil(2.51)
	fmt.Println("ceill", ceil1, ", ceil2", ceil2, ", ceil3 =", ceil3)

	var ceil4, ceil5, ceil6 float64 = math.Ceil(-2.49), math.Ceil(-2.5), math.Ceil(-2.51)
	fmt.Println("ceil =", ceil4, ", ceil5", ceil5, ", ceil6", ceil6)

	// Floor - всегда ↓
	var floor1, floor2, floor3 float64 = math.Floor(2.49), math.Floor(2.5), math.Floor(2.51)
	fmt.Println("floor1", floor1, ", floor2", floor2, ", floor3 =", floor3)

	var floor4, floor5, floor6 float64 = math.Floor(-2.49), math.Floor(-2.5), math.Floor(-2.51)
	fmt.Println("floor4", floor4, ", floor5", floor5, ", floor6 =", floor6)

	// Round 1 по правилам, где >= *.5 округляет дальше от нуля
	var round1, round2, round3 float64 = math.Round(2.49), math.Round(2.5), math.Round(2.51)
	fmt.Println("round1=", round1, ", round2", round2, ", round3", round3)

	var round4, round5, round6 float64 = math.Round(-2.49), math.Round(-2.5), math.Round(-2.51)
	fmt.Println("round4 =", round4, ", round5", round5, ", round6", round6)

	// RoundToEven округление, где *.5 округляет ближайшего чётного числа
	var roundToEven1, roundToEven2 float64 = math.RoundToEven(2.49), math.RoundToEven(2.51)
	fmt.Println("roundToEven1 =", roundToEven1, ", round To Even2 =", roundToEven2)

	var roundToEven3, roundToEven4 float64 = math.RoundToEven(-2.49), math.RoundToEven(-2.51)
	fmt.Println("roundToEven3", roundToEven3, ", round To Even4", roundToEven4)

	var roundToEven5, roundToEven6 float64 = math.RoundToEven(2.5), math.RoundToEven(3.5)
	fmt.Println("roundToEven5", roundToEven5, ", round To Even6", roundToEven6)

	var roundToEven7, roundToEven8 float64 = math.RoundToEven(-2.5), math.RoundToEven(-3.5)
	fmt.Println("roundToEven7", roundToEven7, ", roundToEven8", roundToEven8)

	// Округление до № знаков
	number := 1.45 // -> 1.5
	rounded1 := math.Round(number*10) / 10
	// 1.45 * 10 = 14.5
	// Round (14.5) = 15
	// 15/10 = 1.5
	fmt.Println("rounded1 =", rounded1)

	number = 1.45 // -> 1.4
	rounded2 := math.RoundToEven(number*10) / 10
	// 1.45 10 14.5
	// RoundToEven (14.5) 14
	// 14/10 1.4
	fmt.Println("rounded2 = ", rounded2)

	number = 2.0081 // -> 2.008
	rounded3 := math.Floor(number*1000) / 1000
	// 2.0081 * 1000 2008.1
	// Floor (2008.1) = 2008
	// 2008 / 1000 = 2.008
	fmt.Println(rounded3)

	number = 2.0081 // -> 2.009
	rounded4 := math.Ceil(number*1000) / 1000
	// 2.0081 * 1000 2008.1
	// Ceil(2008.1) = 2009
	// 2009/1000 = 2.009
	fmt.Println(rounded4)

}

func demonstratePow() { // работает с float64
	// Pow
	var powValue1 float64 = math.Pow(1.1, math.Pi) // 1.1 ³
	fmt.Println("powValue1 =", powValue1)

	var int8Value1, int8Value2 int8 = 2, 3
	var powValue2 float64 = math.Pow(float64(int8Value1), float64(int8Value2))
	fmt.Println("powValue2 =", powValue2)

	var powValue3 float64 = math.Pow(-2, 3)
	fmt.Println("powValue3", powValue3)

	// Pow10
	var pow10Value1 float64 = math.Pow10(1)
	fmt.Println("pow10Value1", pow10Value1)

	var pow10Value2 float64 = math.Pow10(2)
	fmt.Println("pow10Value2", pow10Value2)

	// = math. Pow10(2.999)
	var float64Value float64 = 2.999
	var pow10Value3 float64 = math.Pow10(int(float64Value))
	fmt.Println("pow10Value3 =", pow10Value3)
}

func demonstrateMod() {
	// % для целых чисел
	fmt.Println("11/3", 11.0/3.0)
	fmt.Println("11% 3", 11%3)
	for i := 0; i < 3; i++ {
		fmt.Println("Индекс =", 1)
		if i%2 == 0 {
			fmt.Printf("%d чётное число\n", i)
		}
		if i%3 == 0 {
			fmt.Printf("%d делится без остатка на 3\n", i)
		}
	}
}
func runMath() {
	// import "math"
	demonstrateMaxMinForTypes()
	demonstrateMaxMinBetweenValues()
	demonstrateRounding()
	demonstratePow()
	demonstrateMod()
}

// Errors
func demonstrateCheckedUncheckedErrors() {

	int8Value1, err := strconv.ParseInt("127", 10, 8)
	fmt.Printf("int8Value1 = %d, err = %v \n", int8Value1, err)

	int8Value2, err := strconv.ParseInt("127", 10, 8)
	fmt.Printf("int8Value2 = %d, err = %v \n", int8Value2, err)

	fmt.Printf("err as...\n* v: [%v]\n +v: [%+v]\n s: [%s]\n * +s: [%s]\n", err, err, err, err)

	intArray := [2]int{100, 500}
	outOfRangeIndex := len(intArray)
	fmt.Println(intArray[outOfRangeIndex])

	var strErr string = err.Error()
	// var strError string = err
	fmt.Println("err.Error:", strErr)
}

func getPrice(itemInfo string) (float64, error) {
	fmt.Printf("ОБРАБОТКА СТРОКИ [%s]\n", itemInfo)

	if itemInfo == "" {
		return 0, &EmptyItemInfoError{}
	}

	priceIndex := strings.Index(itemInfo, "price:")
	if priceIndex == -1 {
		return 0, &NoPriceInItemInfoError{itemInfo}
	}

	var priceString string = itemInfo[priceIndex+6:]
	priceFloat64, err := strconv.ParseFloat(priceString, 64)
	if err == nil {
		return priceFloat64, nil
	}
	return 0, fmt.Errorf("Некорректная стоимость товара в строке [%s]", itemInfo)
}

//
//func pretendReadingItemInfoFromKeyboard() [] string {
//}
//

func inputNewItemInfo(itemsInfo *[]string) {
}

// Read & Write To File
func writeTextToFile(fileName, str string) {
	/*
				Permissions : для файла

			вот из этой строки, которая ниже os.WriteFile(fileName, bytesToWrite, 0644)
				0 РЕЖИМ ФАЙЛА
				0 - обычный файл
				1 - папка
				2 ссылка
				3 сокет
				4 файл устройства
				5 именованный канал
				6 и 7 зарезервированы для использования в будущем

				ПРАВА ДОСТУПА К ФАЙЛУ
		        0 - это первая цифра (режим файла 👆) 0
				I цифра владелец, права для меня как владельца  0  - нет прав 👇 6
				ІІ цифра группа 4
				ІІІ цифра остальные 4

				0 - Нет прав
				1- Выполнение
				2- Запись
				3 Запись + выполнение
				4 Чтение
				5 Чтение + выполнение
				6- Чтение + запись
				7- Чтение + запись + выполнение
	*/
	//bytesToWrite := []byte(str)     // можно так и так bytesToWrite или []byte(str) ниже вставить
	if err := os.WriteFile(fileName, []byte(str), 0644); err != nil {
		panic(err)
	}
}

func runReadWriteToFile() {
	pricelist := "Mango: - 5.70\nBanana: - 2.35\nOrange: - 2.20\nApple \"Golden\": - 1.95"
	writeTextToFile(TEXT_FILE_NAME, pricelist)
	//defer os.Remove(TEXT_FILE_NAME)  // удаляет файл из папки
	textFromFile := readTextFromFile(TEXT_FILE_NAME)
	fmt.Println(textFromFile)

	documents := []Document{
		{
			Id:   75951,
			Type: "Реализация товаров и услуг",
			Goods: []Good{
				{Id: 75, Name: "Книга Сага (Тонино Бенаквиста, 1997)"},
				{Id: 11, Name: "Доставка"},
			},
			Author: &Employee{Id: 24, Name: "Одри Хорн", Position: "Менеджер по продажам"},
		},
		{
			Id:   13500,
			Type: "Поступление товаров и услуг",
			Goods: []Good{
				{Id: 285, Name: "Кофемашина Samsung"},
				{Id: 289, Name: "Капсулы для кофемашины (500 шт)"},
			},
			Author: &Employee{Id: 7, Name: "Шелли Джонсон", Position: "Менеджер по закупкам"},
		},
	}

	writeDocumentsToJsonFile(JSON_FILE_NAME, &documents)
	//defer os.Remove(JSON_FILE_NAME)  // удаляет файл из папки

	documentsFromFile := readDocumentsFromJsonFile(JSON_FILE_NAME)
	printDocuments("Прочитаны документы:", documentsFromFile)
}

func writeDocumentsToJsonFile(fileName string, documents *[]Document) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(documents)

	printDocuments("Записаны документы:", documents)
}

func printDocuments(header string, documents *[]Document) {
	fmt.Println(strings.ToUpper(header))
	for _, doc := range *documents {
		doc.Print()
	}
}

func readDocumentsFromJsonFile(fileName string) *[]Document {
	bytesFromFile, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	strFromFile := string(bytesFromFile)
	reader := strings.NewReader(strFromFile)
	decoder := json.NewDecoder(reader)

	var documentsFromFile []Document
	err = decoder.Decode(&documentsFromFile)
	if err != nil {
		panic(err)
	}
	return &documentsFromFile
}

func readTextFromFile(fileName string) string {
	str, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(str)
}

// Read From Console
func runReadFromConsole() {

	// import fmt
	fmt.Println("Scan-1. Введите ответ на главный вопрос жизни, вселенной и всего такого...")

	var answer string
	fmt.Scan(&answer)
	fmt.Printf("Ответ на главный вопрос жизни, вселенной и всего такого: [%s]\n\n", answer)

	pause()

	fmt.Println("Scan-2. Это значение ввести не удастся...")
	var skippedVariable string
	fmt.Scan(skippedVariable)
	fmt.Printf("Я же говорю, не удастся: [%s]\n\n", skippedVariable)

	pause()

	fmt.Println("Scan-3. Введите цвет, целое число, дробное число и значение true или false через пробел")
	color, intValue, floatValue, boolValue := "", 0, 0.0, false
	count, err := fmt.Scan(&color, &intValue, &floatValue, &boolValue)
	fmt.Printf("Считано % строк. Ошибка %v\n", count, err)
	fmt.Printf("Цвет: %5, целое число: %d, дробное число: %f, булево: %v\n\n", color, intValue, floatValue, boolValue)

	pause()

	fmt.Println("Scanln-4. Введите любимый фрукт, суммарное количество пальцев на руках...")
	fruit, fingers := "", 0
	count, err = fmt.Scan(&fruit, &fingers)
	fmt.Printf("Считано %d строк. Ошибка = %v\n", count, err)
	fmt.Printf("Любимый фрукт: %5, количество пальцев на руках: %d\n\n", fruit, fingers)

	pause()

	fmt.Println("Scanln-5. Попробуйте ввести топ-3 любимых книги...")
	var favoriteBook1, favoriteBook2, favoriteBook3 string
	fmt.Scanln(&favoriteBook1, &favoriteBook2, &favoriteBook3)
	fmt.Printf("Ваши любимые книги: \n1) [%s]\n2) [%s]\n3) [%s]\n\n", favoriteBook1, favoriteBook2, favoriteBook3)

	fmt.Println("Scan 6. Ещё раз попробуйте ввести топ-3 любимых книги...")
	fmt.Scan(&favoriteBook1, &favoriteBook2, &favoriteBook3)
	fmt.Printf("Ваши любимые книги: \n1) [%s]\n2) [%s]\n3) [%s]\n\n", favoriteBook1, favoriteBook2, favoriteBook3)

	pause()

	// import bufio
	fmt.Println("bufio-7. Ещё раз попробуйте ввести топ-3 любимых книги...")
	reader := bufio.NewReader(os.Stdin)
	favoriteBook1, _ = reader.ReadString('\n')
	favoriteBook2, _ = reader.ReadString('\n')
	favoriteBook3, _ = reader.ReadString('\n')
	fmt.Printf("Ваши любимые книги: \n1) [%s]\n2) [%s]\n3) [%s]\n\n", favoriteBook1, strings.ToUpper(favoriteBook2),
		strings.TrimSpace(favoriteBook3))

	fmt.Println("bufio-8. Введите любимых блогеров, в конце набрав \"конец\"...")
	scanner := bufio.NewScanner(os.Stdin)
	blogers := []string{}
	for {
		scanner.Scan()
		scannedText := scanner.Text()
		if scannedText == "конец" || scannedText == "Конец" {
			break
		}
		blogers = append(blogers, scannedText)
	}
	fmt.Println("Ваши любимые блогеры:")

	for index, bloger := range blogers {
		fmt.Println(index, bloger)
	}
	pause()

}

func pause() {

}
