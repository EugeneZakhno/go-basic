package main

import "fmt"

func main() {
	fmt.Println("Lesson_2:")
	runLessons()
}

func runLesson(fn func()) {
	fn()
}

func runLessons() {
	runLesson(runRunes)
	runLesson(runConstants)
	runLesson(runComments)
	runLesson(runMultipleVars)
	runLesson(runArrays)
	runLesson(runSlices)
	runLesson(runSkipVar)
	runLesson(runIfElseSwitch)
	runLesson(runLoops)
}

// Runes
func printRuneVariables(variableName string, variableValue rune) {
	fmt.Printf("%s as an int: [%d], as a string: [%s] \n", variableName, variableValue, string(variableValue))
}

// Rune - Это полный синоним int32
func runRunes() {
	var runeLetterA rune = 65
	printRuneVariables("Letter 'A'", runeLetterA)
	runeLetterB := 'B'
	printRuneVariables("Letter 'B'", runeLetterB)
	runeLetterC := rune(67)
	printRuneVariables("Letter 'C'", runeLetterC)
	var int32LetterD int32 = 68
	printRuneVariables("Letter 'D'", int32LetterD)
	int32LetterE := int32(69)
	printRuneVariables("Letter 'E'", int32LetterE)
	var runeFireEmoji = '\U0001F525' //
	printRuneVariables("Emoji 'Fire'", runeFireEmoji)

	var stringEmoji = string(128293)
	fmt.Println("stringEmoji =", stringEmoji)
	runeEuroSymbol1 := '\U000020AC' // €
	printRuneVariables("Symbol 'Euro 1'", runeEuroSymbol1)
	runeEuroSymbol2 := '€'
	printRuneVariables("Symbol 'Euro 2'", runeEuroSymbol2)
	var int32NewLineSymbol = '\n'
	printRuneVariables("Symbol 'New Line'", int32NewLineSymbol)
	var runeHammerEmoji = '\U0001F528' //
	printRuneVariables("Emoji 'Hammer'", runeHammerEmoji)
	newEmojiFromFire := runeFireEmoji + 3
	printRuneVariables("New emoji 'Fire+3'", newEmojiFromFire)

}

// Constants
func constFunc() {

}

const NUMBER_OF_WORKERS int8 = 10
const COINS = 100_000
const OSLO_FOUNDATION_YEAR = 1048
const FAVORITE_DRINK = "green tea"
const SMILING_EMOJI = 128512

func runConstants() {
	// FAVORITE_DRINK = "green tea"
	const CONST_FROM_FUNC = "I'm from runConstants() func"
	fmt.Printf("Workers: %d, coins: %d, Oslo foundation year: %d, favorite drink: %s, emoji(d): %d, local const: %s, emoji(s): %s",
		NUMBER_OF_WORKERS, COINS, OSLO_FOUNDATION_YEAR, FAVORITE_DRINK, SMILING_EMOJI, CONST_FROM_FUNC, string(SMILING_EMOJI))
	for i := 0; i < 2; i++ {
		const CONST_FROM_LOOP = "I'm from loop"
		fmt.Println(CONST_FROM_LOOP)
	}
	// fmt.Println(CONST_FROM_LOOP)
	constFunc()
}

// Comments
func runComments() {
}

// MultipleVars
func getThreeVariables() (int32Var1, int32Var2 int32, float32Var float32) {
	int32Var1, int32Var2, float32Var = 9, 10, 11.111111
	return
}
func runMultipleVars() {
	var int8Value1, int8Value2, int8Value3 int8
	int8Value1, int8Value2, int8Value3 = 1, 2, 3
	fmt.Println(int8Value1, int8Value2, int8Value3)
	// var val1, val2 int8, var val3 int = 1, 2, 3
	var intValue, stringValue, runeValue = 4, "five", rune(6)
	fmt.Println(intValue, stringValue, runeValue)
	boolValue, uint32Value := false, uint32(8)
	fmt.Println(boolValue, uint32Value)
	int32Value1, int32Value2, float32Value := getThreeVariables()
	fmt.Println(int32Value1, int32Value2, float32Value)
}

//Arrays - это коллекции однородных данных. Простой - статический

func runArrays() {

	value1, value2 := 1, 2
	fmt.Println("1", value1)
	fmt.Println("2-", value2)

	// 1. var name [size]type
	var stringArray [3]string // ["", "", ""] уже при объявлении под капотом
	fmt.Println(stringArray)

	fmt.Printf("Len = %d, cap = %d\n", len(stringArray), cap(stringArray))
	stringArray[0] = "first"
	fmt.Println(stringArray)
	stringArray[2] = "third"
	fmt.Println(stringArray)
	stringArray[2] = "last"
	fmt.Println(stringArray)
	fmt.Println("[" + stringArray[1] + "]")

	// 2. name := [size] type{value1, value2, value3}
	int8Array := [3]int8{5, 30, 89}
	fmt.Println(int8Array)
	int8Array[0] = 6
	int8Array[2] = 90
	fmt.Println(int8Array)
	// int8Array[3] = 0
	// int8Array[-1] = 0
	// index := 3
	// fmt.Println(int8Array[index])

	// 3. name: [...] type(value, value,..و valueN}
	runeArray := [...]rune{'A', '\U000020AC', 'H', '🔨'}
	runeArray[3] = '€'
	// runeArray [4] = 'B'
	fmt.Printf("Len%d, cap %d\n", len(runeArray), cap(runeArray))

	// 4. Recreation
	fmt.Println("4. Recreation.")
	int16Array := [2]int16{1, 2}
	fmt.Println(int16Array)

	int16Array = [2]int16{3}
	// int 16Array make ([]int16, 2)
	// int16Array = [3]int16(]
	// int16Array = [2]int8{}
	// int16Array= [2]int32{}
	fmt.Println(int16Array)

	// 5. Slices -
	fmt.Println("5. Slices -  в значении обрезать")
	flowersArray := [5]string{"e Daisy", "1 Rose", "2 Violet", "3 Chamomile", "4 Tulip"}
	slice := flowersArray[0:3]
	fmt.Println("slice", slice)

}

// Slices. отличается от массива объявление в виде пустых []
func runSlices() {
	fmt.Println("SLICES - real slices []")
	// 1. name: []type{}
	fmt.Println("1. name: []type{}")
	int8Slice := []int8{}
	fmt.Println(int8Slice)
	fmt.Printf("Length = %d, Capacity = %d\n", len(int8Slice), cap(int8Slice))
	// int8slice[0] = 1
	int8Slice = append(int8Slice, 10)
	int8Slice = append(int8Slice, 20, 30, 40, 50)
	fmt.Println(int8Slice)

	int8Slice = append(int8Slice, int8Slice...)
	int8Slice[0] = 88
	fmt.Println(int8Slice)

	// 2. name:= [] type(vall, val2, ..., valn)
	fmt.Println("// 2. name:= [] type(vall, val2, ..., valn)")
	stringSlice := []string{"cat", "dog"}
	fmt.Println(stringSlice)
	stringSlice = append(stringSlice, []string{"fish", "bird", "deer"}...)
	stringSlice[0], stringSlice[1] = "lama", "crocodile"
	fmt.Println(stringSlice)

	// 3. name: make ([] type, size)
	fmt.Println("3. name: make ([] type, size)")

	runeSlice := make([]rune, 2)
	fmt.Println(runeSlice)
	runeSlice[0] = 'A'
	runeSlice[1] = 'B'
	// runeSlice [2] = 	'C'
	fmt.Println(runeSlice)
	runeSlice = append(runeSlice, 'C', 'D', 'E')
	fmt.Println(runeSlice)
	// 4. len
	intSlice := []int{}
	for i := 0; i < 140; i++ {
		fmt.Printf("Length=%d, Capacity = %d\n", len(intSlice), cap(intSlice))
		intSlice = append(intSlice, i)
	}

	// 5. name := make([]type, len, cap)
	fmt.Println("5. name := make([]type, len, cap)")
	int64Slice := make([]int64, 3, 6)
	fmt.Println(int64Slice)
	fmt.Printf("Length = %d, Capacity = %d\n", len(int64Slice), cap(int64Slice))
	int64Slice[0] = 1
	// int64Slice[3] = 4
	fmt.Println(int64Slice)
	fmt.Printf("Length=%d, Capacity %d\n", len(int64Slice), cap(int64Slice))

	int64Slice = append(int64Slice, []int64{111, 222}...)
	fmt.Println(int64Slice)
	fmt.Printf("Length = %d, Capacity = %d\n", len(int64Slice), cap(int64Slice))

	int64Slice = append(int64Slice, []int64{333, 444}...)
	fmt.Println(int64Slice)
	fmt.Printf("Length = %d, Capacity = %d\n", len(int64Slice), cap(int64Slice))

	// 6. Subslices подсрезы
	fmt.Println("stringSlice:", stringSlice)
	fmt.Println("[0:1):", stringSlice[0:1]) // вторая цифра не вкючительно
	fmt.Println("[0:2):", stringSlice[0:2])
	fmt.Println("[0:0):", stringSlice[0:0])
	fmt.Println("[1:3):", stringSlice[1:3])
	fmt.Println("[:2):", stringSlice[:2]) // это полный аналог 0:2
	fmt.Println("[:1):", stringSlice[:1])
	fmt.Println("[:0):", stringSlice[:0])
	fmt.Println("[0:):", stringSlice[0:]) // это с 0 и до конца (вывести все)
	fmt.Println("[:]:", stringSlice[:])

	// 7. Recreation
	namesSlice := make([]string, 0, 3)
	namesSlice = append(namesSlice, []string{"Dana", "Alice", "Bob"}...)
	fmt.Println(namesSlice)
	fmt.Printf("Length = %d, Capacity = %d\n", len(namesSlice), cap(namesSlice))
	namesSlice = make([]string, 2, 18)
	fmt.Println(namesSlice)
	fmt.Printf("Length = %d, Capacity = %d\n", len(namesSlice), cap(namesSlice))

}

// Skip variables _, _
func demonstrateSkipping(int8Value int8, _ string) (result float32, _ rune) { // _ string - это будет стринг, но она меня не интересует, я это значение даже принимать не буду
	fmt.Println("demonstrateskipping: int8Value =", int8Value)
	return
}

func runSkipVar() {
	float32Result, skippedRune := demonstrateSkipping(1, "2")
	fmt.Println(float32Result, string(skippedRune))
	_, _ = demonstrateSkipping(3, "4")
	demonstrateSkipping(5, "6")
	// demonstrateSkipping (5,)

	for _, value := range []int{8, 90, 350} {
		fmt.Println(value)
	}
}

// If else, switch
func printSeasonOfBirthwithSwitch(monthOfBirth int) {
	switch monthOfBirth {
	case 9, 10, 11: // тоже самое что case 9 || 10 || 11
		fmt.Println("You were born in Autumn")
	case 12, 1, 2:
		fmt.Println("You were born in winter")
		fallthrough // проволиться "сквозь" ровно на 1 блок вниз, т.е. вывести далее (это аналог continue в java)
	case 3, 4, 5:
		fmt.Println("You were born not in Autumn or Summer")
	case 6, 7, 8:
		fmt.Println("You were born in Summer")
	default:
		fmt.Println("Whaaaaaat?!")
	}
}

func runIfElseSwitch() {
	fmt.Println("This line will be printed 100%")
	var intValue int // 0

	if intValue == 0 {
		fmt.Println("This line will be printed because intValue is ZERO")
	}

	intValue = 5

	if intValue == 0 {
		fmt.Println("This line will not be printed")
	} else {
	}
	fmt.Println("This line will be printed because intValue is FIVE now")

	if intValue == 0 {
		fmt.Println("This line will not be printed")
	} else if intValue == 1 {
		fmt.Println("This line will not be printed")
	} else if intValue == 2 {
		fmt.Println("This line will not be printed")
	}

	if intValue == 0 {
		fmt.Println("This line will not be printed")
	} else if intValue == 0 || intValue > 2 { // && - это и , это или
		fmt.Println("This line will be printed because intValue is MORE THAN TWO now")
	} else if intValue == 5 {
		fmt.Println("This line will not be printed")
	} else {
		fmt.Println("This line will not be printed")
	}

	var monthOfBirth string = "October"

	if monthOfBirth == "August" {
		fmt.Println("You were born in August, 8th month, Summer")
	} else if monthOfBirth == "September" {
		fmt.Println("You were born in Septemmer, 9th month, Autumn")
	} else if monthOfBirth == "October" {
		fmt.Println("You were born in October, 10th month, Autumn")
	} else if monthOfBirth == "November" {
		fmt.Println("You were born in November, 11th month, Autumn")
	} else if monthOfBirth == "December" {
		fmt.Println("You were born in December, 12th month, Winter")
	} else if monthOfBirth == "January" {
		fmt.Println("You were born in January, 1th month, Winter")
	} else if monthOfBirth == "February" {
		fmt.Println("You were born in February, 2th month, Winter")
	} else if monthOfBirth == "March" {
		fmt.Println("You were born in March, 3th month, Spring")
	} else if monthOfBirth == "April" {
		fmt.Println("You were born in April, 4th month, Spring")
	} else if monthOfBirth == "May" {
		fmt.Println("You were born in May, 5th month, Spring")
	} else if monthOfBirth == "June" {
		fmt.Println("You were born in June, 6th month, Summer")
	} else if monthOfBirth == "July" {
		fmt.Println("You were born in July, 7th month, Summer")
	}

	// НЕт такого свича, который нельзязаенить if !!!
	switch monthOfBirth {
	case "August":
		fmt.Println("You were born in August, 8th month, Summer")
	case "September":
		fmt.Println("You were born in September, 9th month, Autumn")
	case "October":
		fmt.Println("Switch^ You were born in October, 10th month, Autumn")
	case "November":
		fmt.Println("You were born in November, 11th month, Autumn")
	case "December":
		fmt.Println("You were born in December, 12th month, Winter")
	case "January":
		fmt.Println("You were born in January, 1th month, Winter")
	case "February":
		fmt.Println("You were born in February, 2th month, Winter")
	case "March":
		fmt.Println("You were born in March, 3th month, Spring")
	case "April":
		fmt.Println("You were born in April, 4th month, Spring")
	}
	printSeasonOfBirthwithSwitch(1)
	printSeasonOfBirthwithSwitch(10)
	printSeasonOfBirthwithSwitch(3)
	printSeasonOfBirthwithSwitch(5)
}

// Loops
func runLoops() {
	for i := 0; i < 3; i++ { // i++ - это i = i + 1 или i += 1
		fmt.Println("For i loop #1: index =", i)
	}

	for i := -5; i <= 15; i += 5 {
		fmt.Println("For i loop #2: index =", i)
	}

	/*	for i := 0; i < 4;  { // беспоконечный цикл
			fmt.Println("For i loop #3: index =", i)
			i++
		}
	*/

	/* for i := 0; ; i = 1 { // тоже беспоконечный цикл
		fmt.Println("For i loop #4: index =", i)
		if i == -3 {
			break // прерывает цикл
		}
	}
	*/

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // пропускает итерацию
		}
		fmt.Println("For i loop #5: index =", i)
	}

	/*	for i := 0; i < 10; i++ {
			if i == 5 {
				return // прерывает функцию
			}
		}
	*/

	var index int = 2
	for ; index < 7; index++ {
		if index >= 3 && index <= 5 {
			continue // пропускает итерацию
		}
		fmt.Println("For i loop #5: index", index)
	}
	// goto - переходит к метке

	for {
		fmt.Println("For i Loop #6: index", index)
		if index == 11 {
			break
		}
		index += 2
	}

	/*
		for { //// еще один самый самый Беспоконечный цикл
			fmt.Println("Поставь лайк, напиши коммент, подпишись")
		}
	*/

	for i := 0; i < 10; i++ {
		if i%2 == 0 { // % - остаток от деления/ Делится ли без остатка на 2, на 3 , на 5 и т.д.
			fmt.Println("For i loop #7: EVEN index =", i)
		}
	}

	// Итерация по коллекциям:
	salmonSlice := []string{"I", "am", "a", "singing", "Salmon", "spending", "all", "day", "jamming"}

	fmt.Println("Пронумеровать и пробежать по коллекции:")
	for i, word := range salmonSlice { // range -  ключевое слово range используется для итерации по элементам различных коллекций, таких как массивы, срезы, карты (maps) и каналы.
		fmt.Println(i, word)
	}

	for i, _ := range salmonSlice { // range бежать по коллекции
		fmt.Println(i)
	}

	for _, word := range salmonSlice {
		fmt.Println(word)
	}

	fmt.Println("Вывести только цифры- Output just figures:")
	for i := range salmonSlice { // range бежать по коллекции
		fmt.Println(i)
	}

	for i := 0; i < len(salmonSlice); i += 2 { //i += 2 вывести каждое второе слово
		fmt.Println(i, salmonSlice[i])
	}

	fmt.Println("вывести только слова - output just words:")
	for _, word := range salmonSlice {
		fmt.Println(word)
	}

}
