package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func main() {
	date := Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date.Year)
	date.SetDay(6)
	date.SetMonth(8)
	fmt.Println(date.Month)
	fmt.Println(date.Day)
	fmt.Println(date)
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("Invalid year")
	}
	d.Year = year
	return nil
}
func (d *Date) SetDay(day int) {

}

func (d *Date) SetMonth(month int) {
	d.Month = month
}
