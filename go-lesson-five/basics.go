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
	demonstrateFormatToString()
	demonstrateParse()
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

func demonstrateFormatToString() {
	// Стандарт: 2 января 2006 года 15:04:05 MST-7
	var angelinaJolieBirthDate time.Time
	var locations [100]string
	usaLocation, _ := time.LoadLocation(locations[0])
	angelinaJolieBirthDate = time.Date(1975, time.June, 4, 9, 8, 5, 1, usaLocation) // 4 июня 1975 09:08:05.01
	fmt.Println(angelinaJolieBirthDate)
	var angelinaJolieStrBirthDate string = angelinaJolieBirthDate.Format("02.01.2006 15:04:05")
	fmt.Println(angelinaJolieStrBirthDate)
	fmt.Println(angelinaJolieBirthDate.Format("2.1.06 3:4:5"))
	fmt.Println(angelinaJolieBirthDate.Format("2006-01-02"))
	fmt.Println(angelinaJolieBirthDate.Format("02.01.2006"))
	fmt.Println(angelinaJolieBirthDate.Format("15:04:05"))
	fmt.Println(angelinaJolieBirthDate.Format("Минуты: 04"))
	fmt.Println(angelinaJolieBirthDate.Format("Дата рождения: 02.01.2006 15 часов 04 минут 05 секунд"))
	fmt.Println(angelinaJolieBirthDate.Format("02 Jan 2006"))
	fmt.Println(angelinaJolieBirthDate.Format("02 Feb 2006"))
	fmt.Println(angelinaJolieBirthDate.Format("01/02/2006"))
	fmt.Println(angelinaJolieBirthDate.Format("01/03/2006")) // month -> hours
	fmt.Println(angelinaJolieBirthDate.Format("01/07/2006"))
}

func demonstrateParse() {
	// Стандарт: 2 января 2006 года 15:04:05 MST-7
	var audreyHepburnBirthDate time.Time
	var audreyHepburnBirthDateStr string = "04.05.1929 03:02:01" // 4 Мая 1929 3:02:01
	audreyHepburnBirthDate, _ = time.Parse("02.01.2006 15:04:05", audreyHepburnBirthDateStr)
	fmt.Println("День рождения Одри:", audreyHepburnBirthDate)
	parsedDate, err := time.Parse("День рождения: 02.01", "День рождения: 04.05")
	fmt.Println(parsedDate)
	fmt.Println(err)
	parsedDate, err = time.Parse("День рождения: 02.01", "04.05")
	fmt.Println(parsedDate)
	fmt.Println(err)
	parsedDate, err = time.Parse("15:04:05", " 03:02")
	fmt.Println(parsedDate)
	fmt.Println(err)
	parsedDate, err = time.Parse("15:4", "03:02")
	fmt.Println(parsedDate)
	fmt.Println(err)
	fmt.Println(time.Parse("2006-01-02T15:04:05Z", "1929-05-04T03:02:01Z"))
	fmt.Println(time.Parse("02.01.2006", "04.05.1929"))
	fmt.Println(time.Parse("T15:04:05Z", "T03:02:01Z"))
	fmt.Println(time.Parse("15:04:05", "03:02:01"))
	fmt.Println(time.Parse("15:04", "03:02"))
	fmt.Println(time.Parse("2006-01-02", "1929-05-04"))
	fmt.Println(time.Parse("01/02/2006", "04/05/1929"))
}
