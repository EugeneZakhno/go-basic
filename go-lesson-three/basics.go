package main

import "fmt"

func main() {
	fmt.Println("Lesson_3:")
	runPointer()
}

//func runLessons() {
//	runLesson(runPointer)
//	runLesson(runMap)
//	runLesson(runStructure)
//	runLesson(runPublicPrivate)
//}
//
////Pointer
//
//func printArrayAndPointer(array [3]string, pointer *[3]string) {
//
//}
//
//func changeNames(array [3]string, pointer *[3]string) {
//
//}

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

	fmt.Println("Мои сбережения: %d, твои сбережения: %d \n", mySavings, yourSavings)

}

//Map
