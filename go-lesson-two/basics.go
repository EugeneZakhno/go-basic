package main

func main() {
	a := 4
	println("1. Значение", a)
	println("2.Адрес ", &a)
	changeValueThroughPointer(&a)
	println("3.Значение  ", a)
	squareVal(a)
	println(a)
}
func changeValueThroughPointer(p *int) {
	println(p)
	sqr := *p * *p
	*p = sqr
}

func squareVal(v int) {
	println(v * v)
}
