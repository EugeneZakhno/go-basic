package main //Обязательно main

import (
	"fmt"
	"strconv"
)

/*
ПРИМИТИВНЫЕ ЧИСЛОВЫЕ ТИПЫ ДАННЫХ:
Тип           Бит  Байт       Мин                         Макс
int8           8     1        -128                        127
uint8(byte)    8     1        0                           255
int16         16     2        -32768                      32767
uint16        16     2        0                           65535
int32(rune)   32     4        -2147483648                 2147483647
uint32        32     4        0                           4294967295
int64         64     8        -9223372036854775808        9223372036854775807
uint64        64     8        0                           18446744073709551615
int           ??     ?        ??                          ??
uint          ??     ?        0                           ??

 ЧИСЛА С ПЛАВАЮЩЕЙ ТОЧКОЙ (ЗАПЯТОЙ):
float32       32     4        -3.4e38                     3.4e38
float64       64     8        1.7e308                     1.7e308

 КОМПЛЕКСНЫЕ ЧИСЛА:
complex64     64     8
complex128   128     16
*/

// Global var
var banan string = "Mango"
var favoriteNum int = 100

var int64Value = 100

func main() { //обязательно main
	//var int8Value int8 = 100;
	//var int16Value int16 = -1000
	//var int32Value int32 = 1
	//var int64Value int64 = 1000000
	//var intValue int = 1224444

	var uint8Value uint8 = 111
	//var uint16Value uint16 = 1000
	//var uint32Value uint32 = 1
	var uint64Value uint64 = 1000000
	//var uintValue uint = 1224444

	var float32Value float32 = 1.75221445
	//var float64Value float64 = 1.5

	fmt.Printf("Int value is %d, Int8 value %d, Float Valuue %f \n", uint64Value, uint8Value, float32Value)
	fmt.Printf("Int value = %f \n", float32Value)
	fmt.Printf("Int value = %d \n", uint8Value)

	fmt.Println("Int value = " + strconv.Itoa(int(uint64Value)))

	var favoriteBook string = "Сияние"
	//Local var
	fmt.Println("Hello Go" + "!")
	fmt.Println(favoriteBook)
	favoriteBook = "iota - }{уета это итерация "
	fmt.Println(favoriteBook)
	result := 200
	fmt.Println(result)

	fmt.Println(int64Value)
}
