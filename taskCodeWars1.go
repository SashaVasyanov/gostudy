package divisor

import "fmt"

func main() {
	// OddCount(15000)
	Divisors(30)
}

// func OddCount(n int) int {
// 	var count int
// 	for i := 0; i < n; i++ {
// 		if i%2 != 0 {
// 			fmt.Println(count)
// 			count++
// 		}
// 	}

// 	return count
// }
func Divisors(n int) int {
	var divi []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divi = append(divi, i)
		}
	}
	fmt.Println(divi)
	return len(divi)
}
