package main

import "fmt"

func main() {
	fmt.Println("Lesson_2:")
	runLessons()
}

//func runLesson(fn func()) {
//	fn()
//}

func runLessons() {
	runLesson(runRunes)
	runLesson(runConstants)
	runLesson(runComments)
	runLesson(runMultipleVars)
	runLesson(runArrays)
	runLesson(runSlices)
	runLesson(runSkipVar)
	runLesson(runIfElseSwitch)
	runLesson(Loops)
}

// Runes
func printRuneVariables(variableName string, variableValue rune) {

}

// Constants
func constFunc() {

}

func runConstants() {
}

// Comments
func runComments() {
}

// Это полный синоним int32
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
	var runeFireEmoji rune = '\U0001F525' //
	printRuneVariables("Emoji 'Fire'", runeFireEmoji)
	var stringEmoji string = string(128293)
	fmt.Println("stringEmoji =", stringEmoji)
	runeEuroSymbol1 := '\U000020AC' // €
	printRuneVariables("Symbol 'Euro 1'", runeEuroSymbol1)
	runeEuroSymbol2 := '€'
	printRuneVariables("Symbol 'Euro 2'", runeEuroSymbol2)
	var int32NewLineSymbol int32 = '\n'
	printRuneVariables("Symbol 'New Line'", int32NewLineSymbol)
	var runeHammerEmoji rune = '\U0001F528' //
	printRuneVariables("Emoji 'Hammer'", runeHammerEmoji)
	newEmojiFromFire := runeFireEmoji + 3
	printRuneVariables("New emoji 'Fire+3'", newEmojiFromFire)

}
