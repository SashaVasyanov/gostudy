package main

import "fmt"

type part struct {
	description string
	count       int
}
type schoolBoy struct {
	class string
	name  string
	grade int
}

type sub struct {
	name   string
	rate   float64
	active bool
}
type myNewStruct struct {
	myField int
}

func main() {
	var myStruct struct {
		num   int
		words string
		truth bool
	}
	myStruct.num = 15
	myStruct.words = "Eotova"
	myStruct.truth = false
	fmt.Println(myStruct.num)
	fmt.Println(myStruct.words)
	fmt.Println(myStruct.truth)
	fmt.Printf("%#v\n", myStruct)

	type model struct {
		name  string
		speed float64
	}
	var audi model
	audi.name = "Audi a7"
	audi.speed = 612.12

	var lada model
	lada.name = "Granta"
	lada.speed = 61.24
	fmt.Println("Name:", audi.name, ",", "Speed:", audi.speed)
	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	showInfo(bolts)
	var sasha schoolBoy
	sasha.grade = 5
	sasha.class = "A"
	sasha.name = "Sasha"
	showInf(sasha)

	sub1 := defaultSub("Vadim")
	sub1.rate = 25.2
	printInfo(sub1)
	sub2 := defaultSub("Dasha")
	printInfo(sub2)
	var subWithDiscount sub
	applyDiscount(&subWithDiscount)
	fmt.Println("Rate:", subWithDiscount.rate)
	var value myNewStruct
	value.myField = 5
	var pointer *myNewStruct = &value
	fmt.Println((*pointer).myField)
	//Литералы структур
	newCar := model{name: "Wolk", speed: 571.21}
	fmt.Println("Name:", newCar)
}

func showInfo(par part) {
	fmt.Println("Description:", par.description)
	fmt.Println("Count:", par.count)
}

func showInf(inf schoolBoy) {
	fmt.Println("Name:", inf.name)
	fmt.Println("Grade:", inf.grade)
	fmt.Println("Class:", inf.grade)
}

func printInfo(s sub) {
	fmt.Println("Name:", s.name)
	fmt.Println("Rate:", s.rate)
	fmt.Println("Actie:", s.active)
}
func defaultSub(name string) sub {
	var s sub
	s.name = name
	s.rate = 2.21
	s.active = true
	return s
}
func applyDiscount(s *sub) {
	s.rate = 4.99

}
