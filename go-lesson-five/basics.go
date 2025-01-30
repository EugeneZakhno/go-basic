package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/user"
	"strconv"
	"strings"
	"sync"
	"time"
)

var locations = []string{}

type Song struct {
}

func runLessons() {
	//runLesson(runTime, "Управление временем и простра... продолжительностью")
	//runLesson(runInterfaces, "Интерфейсы: впихнуть невпихуемое")
	//runLesson(runTypeAssertion, "Приведение типов: а ты точно олень?")
	//runLesson(runLog, "Логирование: +10 к уважению на собеседовании")
	//runLesson(runTesting, "Unit тесты с самой популярной библиотекой на GitHub")
	runLesson(runGoroutines, "Горутины: приятнейшая реализация многопоточности")
	//runLesson(runChannels, "Каналы: куда чего послать?")
}

func main() {
	fmt.Println("Lesson_05")
	runLessons()
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

// Утверждение типов (ниже)
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
	if emptyError, ok := err.(*EmptyItemInfoError); ok { // ok - говорит о том, удалось ли нам утвердить этот тип
		fmt.Println(emptyError.Error())
	} else if convertError, ok := err.(*ConvertToInt8Error); ok {
		fmt.Println(convertError.Error())
	}

	// создадим слайс Животные
	animals := []Animal{Racoon{"Гоша"}, Cuckoo{"Зоя"}}
	animalDeer := Animal(Deer{"Олеша"})
	animals = append(animals, animalDeer)

	for _, animal := range animals {

		switch exactAnimal := animal.(type) {
		case Racoon:
			fmt.Println("Я енот", exactAnimal.GetName())
			exactAnimal.InvestigateGarbage()
		case Cuckoo:
			fmt.Printf("Я кукушка %5, а тебе, возможно, осталось %0 лет. Но это не точно. \n", exactAnimal.GetName(), exactAnimal.GetYearLeft())
		case Deer:
			fmt.Printf("Я олень %5 и вот моё фото: %s\n", exactAnimal.name, exactAnimal.GetPicture())
			// case int8:
			// fmt.Println("Невозможно")
		}
	}

	stuff := []interface{}{
		Racoon{"Гоша"},
		int8(15),
	}

	for _, s := range stuff {

		switch e := s.(type) {
		case Racoon:
			fmt.Println("Я енот", e.GetName())
			e.InvestigateGarbage()
		case Cuckoo:
			fmt.Printf("Я кукушка %, а тебе, возможно, осталось %0 лет. Но это не точно. \n", e.GetName(), e.GetYearLeft())
		case Deer:
			fmt.Printf("Я олень %5 и вот моё фото: %s.\n", e.name, e.GetPicture())
		case int8:
			fmt.Println("А я число", e, "типа int8") // e не знаю почему не находит
		}
	}

}

// Логирование
func runLog() {
	// import "log"
	log.Println("func runLog() was called")
	currentUser, _ := user.Current()
	log.Printf("func runLog() was called by user [%s]\n", currentUser.Username)
	log.SetPrefix("Logged record ")
	log.Println("Message with prefix")
	log.SetPrefix("")
	log.Println("Message without prefix")

	filePath := fmt.Sprintf("logs/log_%s.txt", time.Now().Format("2006-01-02 15:04:05"))
	// filePath = ""
	if logFile, err := os.Create(filePath); err != nil {
		log.Panicf("Cannot create file %s", filePath)
	} else {
		defer os.Remove(filePath) // Не перепечатывай эту строку!
		defer logFile.Close()     // Порядок важен.
		log.SetOutput(logFile)
	}

	log.Println("You can't see this message in the debug console")
	log.Println("But you can see it in", filePath)
	log.SetOutput(os.Stdout)
	log.Println("You can't see this message in the file")

	// log.SetFlags (1)
	log.SetFlags(log.Ldate)
	log.Println("Log with dat only (without time)")

	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("Log with date and time")

	log.SetFlags(log.Ltime)
	log.Println("Log with time only (without date)")

	log.SetFlags(log.LstdFlags)
	log.Println("Log with date and time")

	log.SetFlags(log.LstdFlags | log.LUTC)
	log.Println("Log with date and time in UTC, not local time zone")
	// log.Fatalln("\n[ Directed by ) \n [ ROBERT B. WEIDE]")
	// log.Panicln("After this line is printed, panic will begin.")
	log.SetFlags(log.LstdFlags | log.Lmsgprefix)
	log.SetPrefix("msg: ")
	log.Println("Log with prefix after time")
}

func runTesting() {
	//testing in other file youtube-channel.go and youtube-channel_test.go
}

type Resume struct {
	Language          string
	YearsOfExperience int8
	YearsOfEducation  int8
}
type Vacancy struct {
	Language          string
	Company           string
	YearsOfExperience int8
	YearsOfEducation  int8
}

func think() {
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
}

func (v Vacancy) Apply(resume *Resume) (interviewInvitationSent bool) {
	fmt.Printf("%sѕ рассматрирает Ваше резюме на позицию %s\n", v.Company, v.Language)
	think()
	if resume.Language != v.Language {
		fmt.Printf("%5 отклоняет Ваше резюме: вакансия \"%s\" не соответствует \"%s\", которую ищете Вы. \n", v.Company, v.Language, resume.Language)
		return
	}
	think()

	if resume.YearsOfExperience < v.YearsOfExperience {
		fmt.Printf("%s отклоняет Ваше резюме: требуется %0 лет опыта, и Bac %d\n", v.Company, v.YearsOfExperience, resume.YearsOfExperience)
		return
	}
	fmt.Printf("%s убедилась, что у Вас достаточно коммерческого опыта. \n", v.Company)

	think()

	if resume.YearsOfEducation < v.YearsOfEducation {
		fmt.Printf("%S отклоняет Ваше резюме: требуется 20 лет образования, у Вас %d.\n", v.Company, v.YearsOfEducation, resume.YearsOfEducation)
		return
	}
	fmt.Printf("%s убедилась, что у Вас достаточный уровень образования.\n", v.Company)
	think()
	return true

}

func applyForVacanciesFn(resume *Resume, vacancies []Vacancy, applyFunc func(resume *Resume, vacancies []Vacancy) (invitedToInterview []Vacancy)) {

	var invitedToInterview []Vacancy
	timeStart := time.Now()
	invitedToInterview = applyFunc(resume, vacancies)
	fmt.Printf("ВАС ПРИГЛАСИЛИ НА СОБЕСЕДОВАНИЕ % КОМПАНИЙ ПО ВАКАНСИЯМ: \n", len(invitedToInterview))
	for index, vacancy := range invitedToInterview {
		fmt.Printf("%d) %+v \n", index+1, vacancy)
	}
	fmt.Println(strings.Repeat("\n", 2))
	fmt.Println("ПРОДОЛЖИТЕЛЬНОСТЬ В 1 ПОТОКЕ:", time.Now().Sub(timeStart))
	fmt.Println(strings.Repeat("\n", 3))
}

func runGoroutines() {
	goroutinesIntro()
	resume := &Resume{
		Language:          "Golang",
		YearsOfExperience: 2,
		YearsOfEducation:  5,
	}

	vacancies := []Vacancy{
		{Language: "Golang", Company: "Google", YearsOfExperience: 5, YearsOfEducation: 5},
		{Language: "Java", Company: "New Yorker", YearsOfExperience: 2, YearsOfEducation: 4},
		{Language: "Golang", Company: "Microsoft", YearsOfExperience: 0, YearsOfEducation: 5},
		{Language: "C#", Company: "Microsoft", YearsOfExperience: 2, YearsOfEducation: 4},
		{Language: "Golang", Company: "Tesla", YearsOfExperience: 1, YearsOfEducation: 4},
		{Language: "Java", Company: "Google", YearsOfExperience: 5, YearsOfEducation: 0},
		{Language: "Golang", Company: "Sun Microsystems", YearsOfExperience: 2, YearsOfEducation: 4},
		{Language: "Golang", Company: "Booking", YearsOfExperience: 2, YearsOfEducation: 5},
		{Language: "Golang", Company: "Bon Aqua", YearsOfExperience: 0, YearsOfEducation: 5},
		{Language: "Java", Company: "Tesla", YearsOfExperience: 1, YearsOfEducation: 4},
		{Language: "C#", Company: "Tesla", YearsOfExperience: 5, YearsOfEducation: 0},
		{Language: "Golang", Company: "Ugly Coyote", YearsOfExperience: 3, YearsOfEducation: 0},
		{Language: "Golang", Company: "Nike", YearsOfExperience: 3, YearsOfEducation: 4},
		{Language: "C#", Company: "Google", YearsOfExperience: 0, YearsOfEducation: 5},
		{Language: "Golang", Company: "Adidas", YearsOfExperience: 1, YearsOfEducation: 4},
		{Language: "Golang", Company: "Puma", YearsOfExperience: 1, YearsOfEducation: 5},
		{Language: "Java", Company: "Microsoft", YearsOfExperience: 0, YearsOfEducation: 5},
		{Language: "Golang", Company: "Terranova", YearsOfExperience: 2, YearsOfEducation: 4},
		{Language: "Golang", Company: "New Yorker", YearsOfExperience: 0, YearsOfEducation: 4},
	}

	applyForVacanciesFn(resume, vacancies, applyForVacanciesIn1Thread)
	applyForVacanciesFn(resume, vacancies, applyForVacanciesInMultithread)
	applyForVacanciesFn(resume, vacancies, applyForVacanciesInMultithreadWithDoubling)

}

// Многопоточность: Горутины
func goroutinesIntro() {

	var total int
	for i := 1; i <= 100; i++ {
		ind, tot := i, total
		total += i
		fmt.Printf("%d%d%d", ind, tot, total)
	}
	fmt.Println("\nTOTAL =", total)
	fmt.Println("Step here")
	fmt.Println()

	total = 0
	for i := 1; i <= 100; i++ {
		go func(index int) {
			ind, tot := index, total
			total += index
			fmt.Printf("%d%d%d", ind, tot, total)
		}(1)
	}
	fmt.Println("\nTOTAL =", total)
	fmt.Println("Add breakpoint here!")
	fmt.Println("Step here")
	fmt.Println("Step here")
	fmt.Println("TOTAL =", total)
	fmt.Println("Step here")
	fmt.Println()

	// import "sync"
	var wg sync.WaitGroup
	total, maxInt := 0, 100
	// wg.Add(maxInt)
	for i := 1; i <= maxInt; i++ {
		wg.Add(1)
		go func(index int) {
			ind, tot := index, total
			total += index
			fmt.Printf("%d%d%d", ind, tot, total)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("\nTOTAL =", total)
	fmt.Println("Add breakpoint here!")
	fmt.Println("Step here")
	fmt.Println("Step here")
	fmt.Println("TOTAL", total)
	fmt.Println("Step here")
	fmt.Println()
}

// Однопоточный поиск
func applyForVacanciesIn1Thread(resume *Resume, vacancies []Vacancy) (invitedToInterview []Vacancy) {
	for index, vacancy := range vacancies {
		interviewInvitationReceived := vacancy.Apply(resume)
		if interviewInvitationReceived {
			fmt.Printf(" >>> %d: Компания \"%s\" позвала Вас на собеседование на позицию \"%s\",\n", index+1, vacancy.Company, vacancy.Language)
			invitedToInterview = append(invitedToInterview, vacancy)
		} else {
			fmt.Printf(" <<<%d: Компания \"%s\" не готова позвать Вас на собеседование на позицию \"%s\".\n", index+1, vacancy.Company, vacancy.Language)
		}
	}
	return
}

// Как надо делать! Поиск вакансии составил 7 секунд в многопоточной (в однопоточной 56 сек.)
func applyForVacanciesInMultithread(resume *Resume, vacancies []Vacancy) (invitedToInterview []Vacancy) {
	var wg sync.WaitGroup
	wg.Add(len(vacancies))

	for index := range vacancies {
		// go interviewInvitationReceived:vacancy Apply(resume)
		go func(index int) {
			vacancy := vacancies[index]
			interviewInvitationReceived := vacancy.Apply(resume)

			if interviewInvitationReceived {
				fmt.Printf(" >>> %d: Компания \"%s\" позвала Вас на собеседование на позицию \"%s\",\n", index+1, vacancy.Company, vacancy.Language)
				invitedToInterview = append(invitedToInterview, vacancy)
			} else {
				fmt.Printf(" <<< %d: Компания \"%s\" не готова позвать Вас на собеседование на позицию \"%s\",\n", index+1, vacancy.Company, vacancy.Language)
			}
			wg.Done()
		}(index)
	}
	wg.Wait()
	return
}

// О том как не надо делать. Пролетел мимо. Никто не позвал на совбез, т.к. указал pointer в коде и никому не откликнулся получается
func applyForVacanciesInMultithreadWithDoubling(resume *Resume, vacancies []Vacancy) (invitedToInterview []Vacancy) {
	var wg sync.WaitGroup
	wg.Add(len(vacancies))
	for _, vacancy := range vacancies {
		go func(vacancy *Vacancy) {
			interviewInvitationReceived := vacancy.Apply(resume)
			if interviewInvitationReceived {
				fmt.Printf(" >>> Компания \" % s\" позвала Вас на собеседование на позицию \"%s\". \n", vacancy.Company, vacancy.Language)
				invitedToInterview = append(invitedToInterview, *vacancy)
			} else {
				fmt.Printf(" <<< Компания \"%s\" не готова позвать Вас на собеседование на позицию \" % s\".\n", vacancy.Company, vacancy.Language)
			}
			wg.Done()
		}(&vacancy) // Не нужно так делать!  Не следует передавать pointers в многопоточной среде
	}
	wg.Wait()
	return
}

// Многопоточность: Каналы
func printChannelInfo(strChannel chan string, msg string) {
	log.Printf("~~~ %s. Количество слов в каналеЖ %d\n", msg, len(strChannel))
}
func readWriteFromChannel(strChannel chan string) {
	words := []string{"Подписка", "Лайк", "Комментарий"}
	fmt.Println("words:", words)

	printChannelInfo(strChannel, fmt.Sprintf("Канал с размером буфера % создан", cap(strChannel)))
	time.Sleep(time.Second * 3)

	// Записываем все значения из среза слов в пустой канал
	for index := range words {
		go func(channel chan string, wordIndex int) {
			word := words[wordIndex]
			channel <- word
			printChannelInfo(channel, fmt.Sprintf(">>> В канал записано слово \"%s\"", word))
		}(strChannel, index)
	}
	time.Sleep(time.Second * 5)

	printChannelInfo(strChannel, "После записи, перед чтением")
	time.Sleep(time.Second * 2)

	// Просим записать ещё одно значение в заполненный канал
	go func(channel chan string, word string) {
		log.Println("Пауза в новой горутине из-за попытки записать ещё одно значение в заполненный канал...")
		channel <- word
		printChannelInfo(strChannel, "Новая горутина запустилась, удалось записать ещё одно значение в освободившуюся ячейку")
	}(strChannel, "Спасибо")
	time.Sleep(time.Second * 10)

	// Читаем все значения из заполненного канала
	for i := 0; i < len(words)+1; i++ {
		go func(channel chan string) {
			printChannelInfo(channel, fmt.Sprint("<<< Слово из канала: ", <-channel))
			time.Sleep(time.Second * 1)
		}(strChannel)
	}
	time.Sleep(time.Second * 10)

	printChannelInfo(strChannel, "После чтения всех слов")
	time.Sleep(time.Second * 3)

	fmt.Println("Stop here")
}

func demonstrateChannels() {
	var strChannelWith3Cells chan string = make(chan string, 3) // канал с буфером, 3 - это размер буфера
	readWriteFromChannel(strChannelWith3Cells)

	fmt.Println("***********************************************")
	var strChannelWithoutCells chan string = make(chan string) // канал без буфера
	readWriteFromChannel(strChannelWithoutCells)
}

func demonstrate1WayChannels() {
	//var producer1WayChannel <-chan *Resume = make(<-chan *Resume, 1)
	//var consumer1WayChannel chan <- *Resume = make(chan<- *Resume, 1)
	//for i:= 0; i < 3; i++ {
	//newResume := &Resume {Language: "Golang", YearsOfExperience: int8 (rand. Intn (10)), YearsOfEducation: int8(rand.Intn(10))}
	//producer1WayChannel <- newResume
	//
	//go func(producer <-chan *Resume, consumer chan <- *Resume) {
	//	resume := <-producer
	//	fmt.Println(resume)
	//	consumer <- resume
	//}(producer1WayChannel, consumer1WayChannel)
	//	fmt.Println("Резюме:", <- consumer1WayChannel) // Зачем-то же значение было записано в канал...
	//}

	var channel1 chan *Resume = make(chan *Resume, 100) // создаем канал1
	var channel2 chan *Resume = make(chan *Resume, 100) //
	for i := 0; i < 3+rand.Intn(20); i++ {
		newResume := &Resume{Language: "Golang", YearsOfExperience: int8(rand.Intn(10)), YearsOfEducation: int8(rand.Intn(10))}
		channel1 <- newResume

		go func(producer <-chan *Resume, consumer chan<- *Resume) {
			resume := <-producer
			fmt.Printf("%+v \n", resume)
			consumer <- resume
			//producer < &Resume() // invalid operation: cannot send to receive-only channel producer (variable of type <-chan *Resume)
			// res:= <-consumer // invalid operation: cannot receive from send-only channel consumer (variable of type chan< *Resume)
		}(channel1, channel2)
		time.Sleep(time.Second * 5)
		close(channel2)

		fmt.Println("Резюме:")
		for resume := range channel2 {
			fmt.Println(resume)
		}
		// channel2 <- &Resume{Language: "Ошибка-а-а-а-а"} // sendon close channel
		fmt.Println("End")
	}
}

func runChannels() {
	log.SetFlags(log.Ltime)
	demonstrateChannels()
	demonstrate1WayChannels()
}

// Main
func cleanScreen() {
}

func pause() {
}

func conc(strs ...string) (resultStr string) {
	return
}

func runLesson(fn func(), header string) {
	fn() // Вызов функции
}
