package main

func main() {
	a := 4
	println("1. Значение", a)
	println("2.Адрес ", &a)
	changeValueThroughPointer(&a)
	println("5.Значение  ", a)
	squareVal(a)
	println(a)
	squareAdd(&a)
}
func changeValueThroughPointer(p *int) {
	println("3. Из другой функции ", p)
	sqr := *p * *p
	*p = sqr
}

func squareAdd(p *int) { // There is a copy of v
	println(p)
}

func squareVal(v int) { // There is a copy of v
	println(v * v)
}
