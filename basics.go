package main //Обязательно main

import (
	"fmt"
)

/*
ПРИМИТИВНЫЕ ЧИСЛОВЫЕ ТИПЫ ДАННЫХ
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

float32       32     4        -3.4e38                     3.4e38
float64       64     8        1.7e308                     1.7e308

complex64     64     8
complex128   128     16
*/

// Global var
var banan string = "Mango"
var favoriteNum int = 100

func main() { //обязательно main
	var favoriteBook string = "Сияние"
	//Local var
	fmt.Println("Hello Go" + "!")
	fmt.Println(favoriteBook)
	favoriteBook = "iota - }{уета это итерация "
	fmt.Println(favoriteBook)
}
