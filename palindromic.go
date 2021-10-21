package palindromic

import (
	"fmt"
)

func main() {
	// var arr = makeRange(10000, 998001)
	// leftNum := 100
	rightNum := 998001
	var palind int = 0
	var targetPolid int
	arrPal := []int{}

	for i := 999; i >= 900; i-- {
		for j := 999; j >= 900; j-- {
			targetPolid = i * j
			if targetPolid == pali(targetPolid) {
				arrPal = append(arrPal, targetPolid)
				fmt.Println(arrPal)

			}
		}
	}
	for palind <= rightNum {
		if rightNum != pali(targetPolid) {
			palind++

		} else if rightNum == pali(targetPolid) {
			fmt.Println(targetPolid)
		}

	}

}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
func pali(num int) int {
	var remainder, temp int
	var reverse int = 0

	temp = num

	for {
		remainder = num % 10
		reverse = reverse*10 + remainder
		num /= 10

		if num == 0 {
			// Break the loop
			break
		}
	}

	if temp == reverse {
		fmt.Printf("%d is a Palindrome\n", temp)
		return temp
	} else {
		return 0
	}

}
