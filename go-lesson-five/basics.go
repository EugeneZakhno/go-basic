package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Lesson_05")
	runTime()
}
func demonstrateDuration() {}

func demonstrateTimeDifference() {
}
func runTime() {
	// import "time"
	demonstrateBasicDateCreation()
	demonstrateTimeDateCreation()
	demonstrateTimeUnixCreation()
	//demonstrateFormatToString()
	//demonstrateParse()
	//	demonstrateComparison()
	//demonstrateDuration()
	demonstrateTimeDifference()
}

// Интерфейсы
func printSomething(index int, i interface{}) {
}
func runInterfaces() {
}

// Утверждение типов
//func parseStringToInt8(value string) (int, error) {
//	return
//}

// Время и продолжительность
func demonstrateBasicDateCreation() {
	// import time
	var currentTime time.Time = time.Now()
	fmt.Println(currentTime, "\n")
	var currentDay int
	var currentMonth time.Month
	var currentYear int
	currentYear, currentMonth, currentDay = currentTime.Date()
	var currentHour, currentMinute, currentSecond int = currentTime.Clock()
	fmt.Printf("%d.%d.%d %d %d %d \n", currentDay, currentMonth, currentYear, currentHour, currentMinute, currentSecond)
	fmt.Printf("%d %s %d year %02d:%02d:%02d \n", currentDay, currentMonth.String(), currentYear, currentHour, currentMinute, currentSecond)

}

func demonstrateTimeDateCreation() {
	var leoDiCaprioBirthLocation *time.Location
	var locations [100]string
	leoDiCaprioBirthLocation, _ = time.LoadLocation(locations[0])
	fmt.Println(leoDiCaprioBirthLocation)
	leoDiCaprioBirthDate := time.Date(1974, time.November, 11, 2, 47, 4, 8, leoDiCaprioBirthLocation) // 11 Ноября 1974

	leoDay, leoMonth, leoYear := leoDiCaprioBirthDate.Day(), leoDiCaprioBirthDate.Month(), leoDiCaprioBirthDate.Year()
	var leoHour, leoMinute, leoSecond int = leoDiCaprioBirthDate.Hour(), leoDiCaprioBirthDate.Minute(), leoDiCaprioBirthDate.Second()
	var leoNanosecond int = leoDiCaprioBirthDate.Nanosecond()
	var leoLocation *time.Location = leoDiCaprioBirthDate.Location()

	fmt.Println("День рождения Леонардо Ди Каприо:", leoDiCaprioBirthDate)

	fmt.Printf("День рождения Леонардо Ди Каприо в %s: %d.%d.%d %d:%:%d.%d \n",
		leoLocation, leoDay, leoMonth, leoYear, leoHour, leoMinute, leoSecond, leoNanosecond)
	fmt.Printf("День рождения Леонардо Ди Каприо в %s: %02d.%02d.%04d %02d:%02d \n",
		leoLocation, leoDay, int(leoMonth), leoYear, leoHour, leoMinute)
	fmt.Printf("leo родился в %0 день %0 года, день недели: %s \n", leoDiCaprioBirthDate.YearDay(), leoDiCaprioBirthDate.Year(), leoDiCaprioBirthDate.Weekday())
}

func demonstrateTimeUnixCreation() {
	fmt.Println("Точка отсчёта", time.UTC, time.Unix(0, 0).In(time.UTC))
	var johnnyDeppBirthDateInSeconds int64 = -207141360 // 9 Июня 1963 8:44
	var johnnyDeppBirthDate time.Time = time.Unix(johnnyDeppBirthDateInSeconds, 0)
	fmt.Println("День рождения Джонни Деппа по Европе:", johnnyDeppBirthDate)
	var locations [100]string
	usaLocation, _ := time.LoadLocation(locations[0])
	johnnyDeppBirthDate = time.Unix(johnnyDeppBirthDateInSeconds, 0).In(usaLocation)
	fmt.Println("День рождения Джонни Деппа по США:", johnnyDeppBirthDate)
	myLocation, _ := time.LoadLocation(locations[1])
	fmt.Println("Точка отсчёта", myLocation, time.Unix(0, 0).In(myLocation))
}

func demonstrateFormatToString() {}
func demonstrateParse()          {}
