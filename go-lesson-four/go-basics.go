package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Lesson_04")
	demonstrateMaxMinForTypes()
}

const ()

type Author struct {
}

func runLessons() {
	//runLesson (runMath)
	//runLesson (runRandom)
	//runLesson (runSlicesSort)
	//runLesson (runDefer)
	//runLesson (runErrors)
	//runLesson (runReadWriteToFile)
	//runLesson (runReadFromConsole)
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

}
