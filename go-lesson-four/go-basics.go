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
