package main //Обязательно main

import (
	"fmt"
	"strconv"
)

/*
ПРИМИТИВНЫЕ ЧИСЛОВЫЕ ТИПЫ ДАННЫХ:
Тип           Бит  Байт       Мин                     Макс
int8           8     1        -128                    127
uint8(byte)    8     1        0                       255
int16         16     2        -32768                  32767
uint16        16     2        0                       65535
int32(rune)   32     4        -2147483648             2147483647
uint32        32     4        0                       4294967295
int64         64     8        -9223372036854775808    9223372036854775807
uint64        64     8        0                       18446744073709551615
int           ??     ?        ??                      ?? (тоже что и int64) берет из архитектуры компьютера, сервера
uint          ??     ?        0                       ?? (тоже что и uint64) берет из архитектуры компьютера, сервера

 ЧИСЛА С ПЛАВАЮЩЕЙ ТОЧКОЙ (ЗАПЯТОЙ):
float32       32     4        -3.4e38                 3.4e38
float64       64     8        1.7e308                 1.7e308

 КОМПЛЕКСНЫЕ ЧИСЛА:
complex64     64     8
complex128   128     16

 ЗНАЧЕНИЯ ПО УМОЛЧАНИЮ:
 1. Числовые типы
• int, int8, int16, int32, int64: 0
• uint, uint8, uint16, uint32, uint64: 0
• float32: 0.0
• float64: 0.0
• complex64: 0 + 0i
• complex128: 0 + 0i

▎2. Булевый тип

• bool: false

▎3. Символьный тип

• rune (alias for int32): 0 (представляет символ 'u0000', нулевой символ)

▎4. Строки

• string: "" (пустая строка)

▎5. Массивы

• Все элементы массива инициализируются значениями по умолчанию для соответствующего типа.

▎6. Срезы

• slice: nil (срез не инициализирован) !!! Выводит просто [] квадратные скобки

▎7. Карты (maps)

• map: nil (карта не инициализирована)

▎8. Каналы

• channel: nil (канал не инициализирован)

▎9. Структуры

• Все поля структуры инициализируются значениями по умолчанию для их типов.

▎Пример

Вот пример, демонстрирующий значения по умолчанию для различных типов данных:

package main

import "fmt"

type MyStruct struct {
    A int
    B string
}

func main() {
    var i int
    var f float64
    var b bool
    var s string
    var arr [3]int
    var slice []int
    var m map[string]int
    var ch chan int
    var str MyStruct

    fmt.Printf("int: %d\n", i)          // 0
    fmt.Printf("float64: %f\n", f)      // 0.0
    fmt.Printf("bool: %t\n", b)         // false
    fmt.Printf("string: '%s'\n", s)     // ''
    fmt.Printf("array: %v\n", arr)      // [0 0 0]
    fmt.Printf("slice: %v\n", slice)    // []
    fmt.Printf("map: %v\n", m)          // <nil>
    fmt.Printf("channel: %v\n", ch)     // <nil>
    fmt.Printf("struct: %+v\n", str)    // {A:0 B:}
}

*/

// Global var
var banana string = "Mango"
var favoriteNum int = 100

var int64Value = 100

func main() { //обязательно main
	//var int8Value int8 = 100;
	//var int16Value int16 = -1000
	//var int32Value int32 = 1
	//var int64Value int64 = 1000000
	//var intValue int = 1224444

	var runSlice = runSlice()
	fmt.Println("runSlice -> ", runSlice)

	var uint8Value uint8 = 111
	//var uint16Value uint16 = 1000
	//var uint32Value uint32 = 1
	var uint64Value uint64 = 1000000
	//var uintValue uint = 1224444

	var float32Value float32 = 1.75221445
	//var float64Value float64 = 1.5

	var strValueWithIntInside string = "123456"

	var intvalue int
	intvalue, _ = strconv.Atoi(strValueWithIntInside)
	fmt.Printf("strValueWithIntInside is %d \n", intvalue)

	var strInt8Value string = "123"
	tempvalue, err := strconv.ParseInt(strInt8Value, 10, 8)
	fmt.Println(err)

	var int8value = int8(tempvalue)
	fmt.Printf("Int value is %d, Int8 value %d, uInt8 value %d, Float Valuue %f \n", uint64Value, int8value, uint8Value, float32Value)
	fmt.Printf("Int value = %f \n", float32Value)
	fmt.Printf("Int value = %d \n", uint8Value)

	fmt.Println("Int value = " + strconv.Itoa(int64Value))

	var appleIsRed bool = 2*2 == 4
	fmt.Println(appleIsRed)
	appleIsRed = false
	fmt.Println(appleIsRed)

	var favoriteBook string = "Сияние"
	//Local var
	fmt.Println("Hello Go" + "!")
	fmt.Println(favoriteBook)
	favoriteBook = "iota - }{уета это итерация "
	fmt.Println(favoriteBook)
	result := 200
	fmt.Println(result)
	fmt.Println(int64Value)
	fmt.Println(getSum(120, 49))

	sum, sub := getSumAndDiff(300, 120)
	fmt.Printf("Sum = %d, sub = %d", sum, sub)
	fmt.Println(runeRun())

}

func getSum(value1, value2 int) int {
	return value1 + value2
}

func getSumAndDiff(value1, value2 int) (sum int, sub int) {
	sum = value1 + value2
	sub = value1 - value2
	return
}
func isMoreThan10(value int) bool {
	return value > 10
}

func runeRun() rune {
	var runeLettersA rune = 66
	return runeLettersA
}

// ~~~~~~~~~~~~~~~~~~Arrays~~~~~~~~~~~~~~~~~~~~~~~~~~
func runArrays() {
	var stringArray [3]string // by default ["","",""]
	fmt.Println(stringArray)

	var stringArrays = [5]string{"Dsr", "Ger", "dbe", "1244"}

	var arr3 = [3]string{"apple", "banana", "cherry"} // Массив из 3 строк
	fmt.Println(arr3)
	var arr1 = [5]int{1, 2, 3, 4, 5} // Массив из 5 элементов
	fmt.Println(arr1)
	var stringArrawy = [3]string{"Eugene", "38", `Ни фига не делает!`}
	fmt.Println(stringArrays)
	fmt.Println(stringArrawy)
}

func runSlice() [5]int {
	//1. template: name :=[]type{}
	int8Slice := [5]int{45, 78, 11}
	fmt.Println(int8Slice)
	return int8Slice
}
