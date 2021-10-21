package task5

import "fmt"

func main() {
	// 1 * 2 * 3 * 2 * 5 * 7 * 2 * 3 * 11 * 13 * 2 * 17 * 19
	// 1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10 * 11 * 12 * 13 * 14 * 15 * 16 * 17 * 18 * 19 * 20
	var needRange = makeRange(1, 20)
	needInt := 1
	for _, value := range needRange {

		if needInt%value > 0 {
			for _, valueNew := range needRange {
				if (needInt*valueNew)%value == 0 {
					needInt *= valueNew
					break
				}
			}

		}
		fmt.Println(needInt)
	}

}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
