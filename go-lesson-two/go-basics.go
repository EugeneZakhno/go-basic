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
	//runLesson(runSlices)
	//runLesson(runSkipVar)
	//runLesson(runIfElseSwitch)
	//	runLesson(Loops)
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
	int16Array := [2]int16{1, 2}
	fmt.Println(int16Array)

	int16Array = [2]int16{3}
	// int 16Array make ([]int16, 2)
	// int16Array = [3]int16(]
	// int16Array = [2]int8{}
	// int16Array= [2]int32{}
	fmt.Println(int16Array)

	// 6. Slices
	flowersArray := [5]string{"e Daisy", "1 Rose", "2 Violet", "3 Chamomile", "4 Tulip"}
	slice := flowersArray[0:3]
	fmt.Println("slice", slice)

}
func runSlices() {

}
