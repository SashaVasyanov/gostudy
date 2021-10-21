package type1

import "fmt"

type Liters float64
type Gallons float64
type Milliliters float64

func main() {
	// var carFuel Gallons
	// var busFuel Liters
	// carFuel = Gallons(10.0)
	// busFuel = Liters(240.0)
	busFuel := Gallons(10.0)
	carFuel := Liters(240.0)
	fmt.Println(busFuel, carFuel)

	trackFuel := Gallons(50.0)
	airFuel := Liters(260.0)
	trackFuel = Gallons(Liters(260.0) * 0.264)
	airFuel = Liters(Gallons(50.0) * 3.785)
	fmt.Printf("Gallons: %0.1f Litters: %0.1f\n", trackFuel, airFuel)
	autoBusFuel := Gallons(1.2)
	motoCarFuel := Liters(4.5)
	autoBusFuel += LiterstoGallons(Liters(40.0))
	motoCarFuel += GalonstoLiters(Gallons(30.0))
	fmt.Printf("AutoBus: %0.1f gallons\n", autoBusFuel)
	fmt.Printf("MotoCar: %0.1f liters\n", motoCarFuel)
}

func GalonstoLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}
func LiterstoGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}
func MillilitersToGallons(m Milliliters) Gallons {
	return Gallons(m * 0.000264)
}
func GallongToMillimeters(g Gallons) Milliliters {
	return Milliliters(g * 3785.41)
}
