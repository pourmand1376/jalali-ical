package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"ramin.tech/cmd/jalai-ical/cmd/internal/ical"
	"ramin.tech/cmd/jalai-ical/cmd/internal/jalalical"
)

var (
	year = flag.Int("year", time.Now().Year(), "Gregorian Year to be converted")
)

func main() {
	flag.Parse()

	if err := Main(*year); err != nil {
		log.Fatal(err)
	}
}

func Main(year int) error {
	fmt.Printf("generate the persian ical file for year %d", year)

	firstDay := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	ical := ical.NewIcal()
	currday := &firstDay
	for year == currday.Year() {
		day := *currday
		jalali := jalalical.NewJalaliCal(day)
		ical.AddEvent(day, jalali.FormatDay())
		day = currday.Add(24 * time.Hour)
		year = currday.Year()
		currday = &day
		log.Println(jalali.FormatDay())
		log.Println(year)
	}
	iclContent := ical.Serialize()
	if iclContent == "" {
		panic("something went wrong")
	}

	firstJalaliYear := jalalical.NewJalaliCal(firstDay)
	secondJalaliYear := jalalical.NewJalaliCal(*currday)
	// persian (jalali) calendar, filename
	fileName := fmt.Sprintf("%d_(%d-%d)_persian_calendar.ics",
		year, firstJalaliYear.Year(), secondJalaliYear.Year())

	if err := writeToFile(fileName, iclContent); err != nil {
		return err
	}

	return generateGarbageCalendars(year)
}

func generateGarbageCalendars(year int) error {
	firstDay := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	tehranLoc := jalalical.TehranLocation()

	oddIcal := ical.NewIcal()
	evenIcal := ical.NewIcal()

	currday := &firstDay
	for year == currday.Year() {
		day := *currday
		tehranDay := day.In(tehranLoc)
		startTime := time.Date(tehranDay.Year(), tehranDay.Month(), tehranDay.Day(), 21, 0, 0, 0, tehranLoc)
		endTime := startTime.Add(10 * time.Minute)

		jalali := jalalical.NewJalaliCal(tehranDay)
		const title = "Garbage collection"
		if jalali.Day() != 31 {
			if jalali.Day()%2 == 1 {
				oddIcal.AddTimedEvent(startTime, endTime, title)
			} else {
				evenIcal.AddTimedEvent(startTime, endTime, title)
			}
		}

		day = currday.Add(24 * time.Hour)
		year = currday.Year()
		currday = &day
	}

	firstJalaliYear := jalalical.NewJalaliCal(firstDay)
	secondJalaliYear := jalalical.NewJalaliCal(*currday)

	oddContent := oddIcal.Serialize()
	evenContent := evenIcal.Serialize()

	oddFileName := fmt.Sprintf("%d_(%d-%d)_garbage_odd_calendar.ics",
		year, firstJalaliYear.Year(), secondJalaliYear.Year())
	evenFileName := fmt.Sprintf("%d_(%d-%d)_garbage_even_calendar.ics",
		year, firstJalaliYear.Year(), secondJalaliYear.Year())

	if err := writeToFile(oddFileName, oddContent); err != nil {
		return err
	}
	return writeToFile(evenFileName, evenContent)
}

func writeToFile(fileName string, data string) error {
	log.Println("writing to file " + fileName)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}
