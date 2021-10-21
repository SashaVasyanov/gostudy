package method

import (
	"fmt"
)

var a int
var b int

type MyType string
type RandNum int

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}
func main() {

	value := MyType("this is value")
	value.sayHi()
	othersValue := MyType("we're others values")
	othersValue.sayHi()
	value1 := MyType("Vasya")
	value1.MethodWithParam(5, true)
	Sasha := MyType("I'm")
	Sasha.LoveOrNot("Sasha", "love", "girl", true)
	length := MyType("I want return length this string")
	fmt.Println(length.retrurnLen())
	amount := 6
	double(&amount)
	fmt.Println(amount)
	newInt := 12
	double(&newInt)
	fmt.Println(newInt)
}

func (m MyType) MethodWithParam(num int, be bool) {
	fmt.Println(m)
	fmt.Println(num)
	fmt.Println(be)
}
func (m MyType) LoveOrNot(name string, love string, girl string, truth bool) {
	fmt.Println(m, name, love, girl)
}
func (m MyType) retrurnLen() int {
	return len(m)
}
func double(number *int) {
	*number *= 2
}
