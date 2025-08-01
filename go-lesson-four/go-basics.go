package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
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
	Name        string
	YearOfBirth int
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

func runLesson(m func()) {
	m() // Вызов функции m
}

// ******************************************************************************************************************************************************************************
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
	var float64Value1, float64Value2 float64 = 10.1, 3.9
	//fmt.Println(float64Valuel % float64Value2) // invalid operation: operator % not defined on float64Valuel (variable of type float64)
	fmt.Println(int(float64Value1) % int(float64Value2))
	fmt.Println(int(float64Value1), int(float64Value2))

	fmt.Println((7 % -4))
	fmt.Println((-7 % 4))

	// Mod для чисел с плавающей точкой
	fmt.Println(math.Mod(9.999, 3.333))  // 9.999/3.333 = 3.0
	fmt.Println(math.Mod(-9.999, 3.333)) // -9.999/3.333 = -3.0
	fmt.Println(math.Mod(9.999, -3.333)) // -9.999/3.333 = -3.0

	intValue1, intValue2 := 13, 5
	result := math.Mod(float64(intValue1), float64(intValue2))
	fmt.Println(result)

	fmt.Println(math.Mod(10, 3))
}

func runMath() {
	// import "math"
	demonstrateMaxMinForTypes()
	demonstrateMaxMinBetweenValues()
	demonstrateRounding()
	demonstratePow()
	demonstrateMod()
}

// ******************************************************************************************************************************************************************************
// Random
func runRandom() {
	// import math/rand
	var arrayFromToMaxInt [3]int
	var arrayFromTo99 [3]int

	for i := 0; i <= 2; i++ {
		var randInt int = rand.Int()
		arrayFromToMaxInt[i] = randInt
		fmt.Println(arrayFromToMaxInt)
		var randIntN int = rand.Intn(100)
		arrayFromTo99[i] = randIntN
		fmt.Println(arrayFromTo99)
	}
	var randInt32 int32 = rand.Int31() // Обнять...
	var randInt64 int64 = rand.Int63() // ... и плакать
	fmt.Println(randInt32, randInt64)

	var randInt32N int32 = rand.Int31n(100) // Принять...
	var randInt64N int64 = rand.Int63n(100) // ...и смириться
	fmt.Println(randInt32N, randInt64N)

	var randUint32 uint32 = rand.Uint32()
	var randUint64 uint64 = rand.Uint64()
	fmt.Println(randUint32, randUint64)

	var randFloat32 float32 = rand.Float32()
	var randFloat64 float64 = rand.Float64()
	fmt.Println(randFloat32, randFloat64)

	// _ = rand.Intn(-1)
	arrayInt8N := [30]int8{}
	for i := 0; i < len(arrayInt8N); i++ {
		// 128 + 127 255 -> [0:255] -> [0:256)
		randInt8InDefaultInt := rand.Intn(256) - 128 // [-128:127]
		arrayInt8N[i] = int8(randInt8InDefaultInt)
	}
	fmt.Println("array Int8N =", arrayInt8N)
	fmt.Println()

	arrayIntBetween10And100 := [30]int{}
	for i := 0; i < len(arrayIntBetween10And100); i++ {
		// [10:100] = [10:101) [0:91) + 10
		randIntNBetweenAnd90 := rand.Intn(91)                  // [0:90]
		arrayIntBetween10And100[1] = 10 + randIntNBetweenAnd90 // [0+10:90+10] [10:100]
		fmt.Println("array IntBetween10And100", arrayIntBetween10And100)
		fmt.Println()
		arrayFloat64Between10And100 := [10]float32{}
		for i := 0; i < len(arrayFloat64Between10And100); i++ {
			randIntNBetweenAnd89 := rand.Intn(90) // [0:90) [0:89]
			randFoat32 := rand.Float32()
			// [0.0,1.0) [0.0,0.(9)]
			arrayFloat64Between10And100[i] = float32(10+randIntNBetweenAnd89) + randFoat32 // [0+0.0: 89+0.(9)] [0.0: 89. (9)]
		}
		fmt.Println("array Float64Between10And100", arrayFloat64Between10And100)
		fmt.Println()
		// До Go 1.20 генератор заполнялся, как Seed(1), при запуске программы.
		// Устарел: начиная с Go 1.20 нет причин вызывать Seed со случайным значением.
		// Если Seed не вызывается, генератор заполняется случайным образом при запуске программы-
		rand.Seed(time.Now().UnixNano())
	}
}

// ******************************************************************************************************************************************************************************
// Slices Sort
func runSlicesSort() {
	// import "sort"
	println("👥 Slices Sort:✏👣 ")
	intSlice := []int{1, 7, 8, -23, 9, -8}
	if !sort.IntsAreSorted(intSlice) { // утверждение = false     утверждение != true...
		sort.Ints(intSlice)
		fmt.Println("sorted intSlice", intSlice)
	}

	intsSlice := []int{5, 3, -8, 1, 2, 10}

	// Bubble sort
	n := len(intsSlice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if intsSlice[j] > intsSlice[j+1] {
				// Меняем местами элементы, если они находятся в неправильном порядке
				intsSlice[j], intsSlice[j+1] = intsSlice[j+1], intsSlice[j]
			}
		}
	}
	fmt.Println("Sorted array:", intsSlice)

	//	reverse with loop (от большему  к меньшему)
	for i, j := 0, len(intSlice)-1; i < j; i, j = i+1, j-1 {
		intSlice[i], intSlice[j] = intSlice[j], intSlice[i]
	}
	fmt.Println("Reverse with loop big to small", intSlice)

	//reverse with loop (от меньшего к большему)
	for i, j := 0, len(intSlice)-1; i < j; i, j = i+1, j-1 {
		intSlice[j], intSlice[i] = intSlice[i], intSlice[j]
	}
	fmt.Println("Reverse with loop from small to big", intSlice)

	stringSlice := []string{"Bbb", "AAA", "A", "D", "Ccc", "Bb"}
	if !sort.StringsAreSorted(stringSlice) {
		sort.Strings(stringSlice)
		fmt.Println("Reverse stringSlice", stringSlice)
	}

	float64Slice := []float64{4.444, 5.555, 1.111, 6.666, 3.333, 2.222}
	if !sort.Float64sAreSorted(float64Slice) {
		sort.Float64s(float64Slice)
		fmt.Println("sorted float64Slice", float64Slice)
	}

	println("\n Анонимная функция : ")
	floatSlice := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6}
	rand.Shuffle(len(floatSlice), func(i, j int) { // анонимная функция с аргументами i - индекс одного числа, ј - индекс другого числа
		floatSlice[i], floatSlice[j] = floatSlice[i]+1, floatSlice[j]+1
	})
	fmt.Println("shuffled floatSlice", floatSlice)

	authors := []Author{{"Стивен Кинг", 1947}, {"Джоан Роулинг", 1965}, {"Нил Гейман", 1960}, {"Макс Фрай", 1965}}
	fmt.Println(authors)

	sort.Slice(authors, func(i, j int) bool {
		if authors[j].YearOfBirth != authors[i].YearOfBirth {
			return authors[j].YearOfBirth > authors[i].YearOfBirth
		} else if len(authors[j].Name) != len(authors[i].Name) {
			return len(authors[j].Name) > len(authors[i].Name)
		}
		return authors[j].Name > authors[i].Name
	})
	fmt.Println(authors)

	compareAnonFunc := func(i, j int) bool {
		fmt.Printf("\nCравниваем %v (j=%d) и %v (i=%d)\n", authors[j], j, authors[i], i)

		// У автора с индексом j больше год рождения, чем у автора с индексом і?
		if authors[j].YearOfBirth > authors[i].YearOfBirth {
			fmt.Printf(" Год рождения %5 (j-%d) больше чем у %s (1-%d). Меняем их местами (true).\n", authors[j].Name, j, authors[i].Name, i)
			fmt.Println("-", authors)
			return true
		} else if authors[j].YearOfBirth < authors[i].YearOfBirth {
			fmt.Printf(" - Год рождения %s (j-%d) меньше чем %s (1-%d). Оставляем их на местах (false).\n", authors[j].Name, j, authors[i].Name, i)
			fmt.Println("-", authors)
			return false
		}
		// У автора с индексом і длиннее имя, чем у автора с индексом і?
		if len(authors[j].Name) > len(authors[i].Name) {
			fmt.Printf(" Имя %s (1-%d) длиннее у %s (1=%d). Меняем их местами (true).\n", authors[j].Name, j, authors[i].Name, i)
			fmt.Println("", authors)
			return true
		} else if len(authors[j].Name) < len(authors[i].Name) {
			fmt.Printf(" Имя %s (j=%d) короче %s (1=%d). Оставляем их на местах (false).\n", authors[j].Name, j, authors[1].Name, i)
			fmt.Println("-", authors)
			return false
		}
		// У автора с индексом j "больше" имя, чем у автора с индексом і?
		return authors[j].Name > authors[i].Name
	}

	sort.Slice(authors, compareAnonFunc)
	fmt.Println(authors)
}

// ******************************************************************************************************************************************************************************
// Defer , кстати return выполнится после defer | А еще LIFO
func demonstrateOneDeferCall() {
	defer fmt.Println("Эта строка выведется под номером", 6, "после выполнения функции деmonstrate OneDeferCall()")
	fmt.Println("Эта строка выведется под номером", 1)
	for i := 2; i <= 4; i++ {
		fmt.Println("Эта строка из цикла выведется под номером", i)
	}
	fmt.Println("Эта строка выведется под номером", 5)
}

func demonstrateFourDeferCalls() {
	defer fmt.Println("Эта строка выведется под номером", 9, "после выполнения функции demonstrate FourDeferCalls()")
	fmt.Println("Эта строка вызовется под номером", 1)
	for i := 2; i < 5; i++ {
		fmt.Println("Эта строка из цикла выведется под номером", 1)
		if i == 3 {
			defer fmt.Println("Эта строка выведется под номером", 8, "после выполнения функции demonstrate FourDeferCalls()")
		} else if i == 4 {
			defer fmt.Println("Эта строка выведется под номером", 7, "после выполнения функции demonstrate FourDeferCalls()")
		}
	}
	fmt.Println("Эта строка выведется под номером", 5)
	defer fmt.Println("Эта строка выведется под номером", 6, "после выполнения функции demonstrate FourDeferCalls()")
}

func runDefer() {
	defer fmt.Println(string('\U0001F609'), "aspdof", string('\U0001F44F'))
	demonstrateOneDeferCall()
	demonstrateFourDeferCalls()
}

// ******************************************************************************************************************************************************************************
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

func inputNewItemInfo(itemsInfo *[]string) {
}

func pretendReadingItemInfoFromKeyboard() []string {
	itemsInfo := []string{
		"item: Kaktus, price: 4.30",
		// Корректная строка
		// Пустая строка
		"item: Sombrero",
		// Строка без цены
		"item: Lime, price: 1.5.0",
		// Строка с некорректной ценой
	}
	fmt.Println("Добавлена информация товарах:")
	for _, itemInfo := range itemsInfo {
		fmt.Printf("[%s]\n", itemInfo)
	}
	return itemsInfo
}

func demonstrateErrors() (totalPrice float64, itemsInfoError error) {
	itemsInfo := pretendReadingItemInfoFromKeyboard()
	for i := 0; ; i++ {
		if i == len(itemsInfo) {
			break
		}
		itemInfo := itemsInfo[i]

		price, err := getPrice(itemInfo)
		if err == nil {
			totalPrice += price
			fmt.Println(" + Общая сумма увеличилась на", price, "и теперь составляет", totalPrice)
			continue
		}
		if itemsInfoError == nil {
			// import Errors
			itemsInfoError = errors.New("Сумма по прайсу не может быть рассчитана корректно, во вермя вводы было допущено не менее 1 ошибки.")
		}

		if emptyItemInfoError, ok := err.(*EmptyItemInfoError); ok {
			fmt.Println(emptyItemInfoError, "Повторите попытку.")
			inputNewItemInfo(&itemsInfo)
		} else if noPriceInItemInfoError, ok := err.(*NoPriceInItemInfoError); ok {
			fmt.Println(noPriceInItemInfoError)
		} else {
			fmt.Println("Другая ошибка:", err.Error())
		}
	}
	return
}

func demonstratePanic() {
	intArray := [2]int{100, 500}
	outOfRangeIndex := len(intArray)
	fmt.Println("Получаю значение по инексу, выходящему за границу массива...")
	fmt.Println(intArray[outOfRangeIndex])
	fmt.Println("Недостижимое продолжение работы функции demonstratePanic()")
}

func demonstratePanicAndRecover() {
	defer func() {
		if arg := recover(); arg != nil {
			fmt.Println("ПАНИКА:", arg)
			fmt.Println("Отставить панику!")
		}
	}()
	demonstratePanic()
	fmt.Println("Недостижимое продолжение работы функции demonstrate PanicAnd Recover()")
}

func runErrors() {

	demonstrateCheckedUncheckedErrors()

	if totalPrice, err := demonstrateErrors(); err == nil {
		fmt.Println("ОБЩАЯ СТОИМОСТЬ:", totalPrice)
	} else {
		fmt.Println(err)
	}
	demonstratePanicAndRecover()
	fmt.Println("Продолжение работы функции runErrors()")
}

// **********************************************************************************************************************************************************************
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

// ******************************************************************************************************************************************************************************
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

//Main

func cleanScreen() {

}

func pause() {

}
