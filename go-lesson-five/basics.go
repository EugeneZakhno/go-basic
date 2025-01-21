package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Lesson_05")
	runTime()
	runInterfaces()
}

func runTime() {
	// import "time"
	demonstrateBasicDateCreation()
	demonstrateTimeDateCreation()
	demonstrateTimeUnixCreation()
	demonstrateFormatToString()
	demonstrateParse()
	demonstrateComparison()
	demonstrateDuration()
	demonstrateTimeDifference()

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

func demonstrateComparison() {
	now := time.Now()
	tokenExpiration := now.Add(time.Hour * 2)
	fmt.Printf("Сейчас %02d:%02d:%02d, срок действия токена заканчивается в %02d:%02d:%02d", now.Hour(),
		now.Minute(), now.Second(), tokenExpiration.Hour(), tokenExpiration.Minute(), tokenExpiration.Second())
	fmt.Println("Токен уже деактивирован?", tokenExpiration.Before(now))
	fmt.Println("Токен ещё активен?", now.Before(tokenExpiration))

	var sameTime time.Time = now.Add(time.Hour * 2)
	fmt.Printf("Срок действия токена [%s] равен таким же дате и времени [%s]? %v \n", tokenExpiration, sameTime,
		tokenExpiration.Equal(sameTime))

	var futureTime time.Time = tokenExpiration.Add(time.Minute * 65)
	fmt.Printf("Cрок действия токена [%s] равен большей дате [%s]? % \n", tokenExpiration, futureTime, tokenExpiration.Equal(futureTime))

	var pastTime time.Time = tokenExpiration.Add(time.Second - 3600)
	fmt.Printf("Cрок действия токена [%s] равен меньшей дате [%s]? %v \n", tokenExpiration, pastTime, tokenExpiration.Equal(pastTime))

	fmt.Println("Compare: сравнить одинаковые значение:", tokenExpiration.Compare(now.Add(time.Hour*2))) // 0 - это true
	fmt.Println("Compare: сравнить меньшее большим:", now.Compare(tokenExpiration))                      // -1
	fmt.Println("Compare: сравнить большее меньшим:", tokenExpiration.Compare(now))                      // 1
}

func demonstrateDuration() {
	var soupCookingStartTime time.Time = time.Now()
	var soupCookingDuration time.Duration = time.Duration(time.Hour*2 + time.Minute*15)
	fmt.Println("Сейчас", soupCookingStartTime)
	fmt.Println("Варить суп следует", soupCookingDuration)

	soupCookingDuration += time.Second * 75

	fmt.Println("Но лучше варить его ровно", soupCookingDuration)

	// soupWillBeReady time.Hour
	// soupWillBeReady soupWill BeReady

	var spoiledSoupDuration time.Duration = soupCookingDuration
	spoiledSoupDuration -= time.Hour
	fmt.Println("Cуп будет недоварен если готовить его", spoiledSoupDuration)
	fmt.Printf("Правильная продолжительность варки супа %, либо % часов, либо % минут, либо % секунд, либо же 1% миллисекунд \n",
		soupCookingDuration, soupCookingDuration.Hours(), soupCookingDuration.Minutes(), soupCookingDuration.Seconds(), soupCookingDuration.Milliseconds())

	var soupWillBeReady time.Time = soupCookingStartTime.Add(soupCookingDuration)
	fmt.Println("Супер должен быть готов к", soupWillBeReady)

	var startToThinkAboutSoupTime time.Time = soupCookingStartTime.Add(time.Hour * -4)
	fmt.Println("Начала думать о приготовлении супа", startToThinkAboutSoupTime)

	var thinkingAboutSoupDuration time.Duration = startToThinkAboutSoupTime.Sub(soupCookingStartTime)
	fmt.Println("От идеи приготовить суп до начала его приготовления прошло", thinkingAboutSoupDuration)
}

func demonstrateTimeDifference() {
	date2023 := time.Date(2023, 12, 31, 23, 59, 59, 0, time.Local)
	date2024 := date2023.Add(time.Duration(time.Second * 2))

	readableFormat := "02.01.2006 15:04:05"
	fmt.Println(date2023.Format(readableFormat))
	fmt.Println(date2024.Format(readableFormat))

	var yearsBetween int = date2024.Year() - date2023.Year()
	fmt.Println("Неправильная разница в годах: ", yearsBetween)

	date2025 := date2023.Add(time.Duration(time.Hour*24*(365+1) + time.Second*2))
	fmt.Println(date2025.Format(readableFormat))

	var differenceInSeconds int64 = date2025.Unix() - date2023.Unix()
	fmt.Println("Разница в годах:", differenceInSeconds/(60*60*24*365))
	fmt.Println("Разница в днях:", differenceInSeconds/(60*60*24))
	fmt.Println("Разница в часax:", differenceInSeconds/(60*60))
	fmt.Println("Разница в минутах:", differenceInSeconds/60)
	fmt.Println("Разница в секундах:", differenceInSeconds)
}

//  Interfaces

func printSomething(index int, i interface{}) {
	fmt.Printf("%d", "Что-то не понятного типа: %+v\n", index+1, i)
}

func runInterfaces() {
	racoon := Racoon{"Гоша"}
	cuckoo := Cuckoo{"Зоя"}

	_, _ = []Racoon{racoon}, []Cuckoo{cuckoo}
	// = []Racoon{racoon, cuckoo), [] Cuckoo{racoon, cuckoo}

	animals := []Animal{racoon, cuckoo}
	animals = append(animals, Deer{"Олеша"})
	// animals append(animals, Animal ("Безымянный зверь"})
	fmt.Printf("%+v \n", animals)

	for _, animal := range animals {
		fmt.Printf("> %s говорит:\n", animal.GetName())
		animal.Print()
	}

	things := []interface{}{
		42,
		"Don't panic!",
		animals,
		errors.New("Какая-то ошибка"),
		nil,
		runInterfaces,
	}
	for i, thing := range things {
		printSomething(i, thing)
	}
}

// Утверждение типов
func parseStringToInt8(value string) (int8, error) {

	if value == "" {
		return 0, &EmptyItemInfoError{}
	}

	parsedValue, err := strconv.ParseInt(value, 10, 8)

	if err != nil {
		return 0, &ConvertToInt8Error{value}
	}

	return int8(parsedValue), nil
}

func runTypeAssertion() {
	_, err := parseStringToInt8("")
	emptyError := err.(*EmptyItemInfoError)
	fmt.Println(emptyError.Error())

	_, err = parseStringToInt8("128")
	if emptyError, ok := err.(*EmptyItemInfoError); ok {
		fmt.Println(emptyError.Error())
	} else if convertError, ok := err.(*ConvertToInt8Error); ok {
		fmt.Println(convertError.Error())
	}

	animals := []Animal{Racoon{"Гоша"}, Cuckoo{"Зоя"}}
	animalDeer := Animal(Deer{"Олеша"})
	animals = append(animals, animalDeer)
}
