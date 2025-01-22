package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Lesson_04")
	//demonstrateMaxMinForTypes()
	//demonstrateCheckedUncheckedErrors()
	runLessons()
	runReadFromConsole()
}

const (
	TEXT_FILE_NAME = "documents/price-list.txt"
	JSON_FILE_NAME = "documents/documents.json"
)

type Author struct {
}

func runLessons() {
	runLesson(runMath)
	runLesson(runRandom)
	runLesson(runSlicesSort)
	runLesson(runDefer)
	runLesson(runErrors)
	runLesson(runReadWriteToFile)
	runLesson(runReadFromConsole)
}

// MAth
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

func demonstrateMaxMinBetweenValues() {
}

func demonstrateRounding() {
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
			 РЕЖИМ ФАЙЛА
			0 - обычный файл
			1 - папка
			2 ссылка
			3 сокет
			4 файл устройства
			5 именованный канал
			6 и 7 зарезервированы для использования в будущем

			ПРАВА ДОСТУПА К ФАЙЛУ
			I цифра владелец, права для меня как владельца  0  - нет прав 👇
			ІІ цифра группа
			ІІІ цифра остальные

			0 - Нет прав
			1- Выполнение
			2- Запись
			3 Запись + выполнение
			4 Чтение
			5 Чтение + выполнение
			6- Чтение + запись
			7- Чтение + запись + выполнение
	*/
	bytesToWrite := []byte(str)
	if err := os.WriteFile(fileName, bytesToWrite, 0644); err != nil {
		panic(err)
	}
}

func runReadWriteToFile() {
	pricelist := "Mango: 5.70\nBanana: 2.35\nOrange: 2.20\nApple \"Golden\":1.95"
	writeTextToFile(TEXT_FILE_NAME, pricelist)
	defer os.Remove(TEXT_FILE_NAME)
	textFromFile := readTextFromFile(TEXT_FILE_NAME)
	fmt.Println(textFromFile)
	documents := []Document{}

	writeDocumentsToJsonFile(JSON_FILE_NAME, &documents)
	defer os.Remove(JSON_FILE_NAME)

	documentsFromFile := readDocumentsFromJsonFile(JSON_FILE_NAME)
	printDocuments("Прочитаны документы:", documentsFromFile)
}

// Read From Console
func runReadFromConsole() {

	// import fmt
	fmt.Println("Scan-1. Введите ответ на главный вопрос жизни, вселенной и всего такого...")

	var answer string
	fmt.Scan(&answer)
	fmt.Printf("Ответ на главный вопрос жизни, вселенной и всего такого: [%s]\n\n", answer)

	pause()
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
		if scannedText == "конец" {
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
