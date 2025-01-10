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

}

//Map
