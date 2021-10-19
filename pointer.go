package main

import "fmt"

type Fmt string
type Liters float64
type Gallons float64
type Milliliters float64

func main() {
	value := Fmt("value")
	pointer := &value
	value.method()
	pointer.pointerMeth()
	soda := Liters(3)
	fmt.Printf("%0.1f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(60)
	fmt.Printf("%0.1f milliters equal %0.3f gallons\n", water, water.ToGallons())
}

func (m Fmt) method() {
	fmt.Println("Method without pointer")
}
func (m *Fmt) pointerMeth() {
	fmt.Println("Method with pointer")
}
func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}
func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}
