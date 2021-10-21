package map1

import (
	"fmt"
	"sort"
)

func main() {
	var ok bool
	rings := make(map[string]int)
	var ring int
	rings["First song"] = 1
	rings["Second song"] = 2
	rings["Third"] = 3
	rings["Fourth"] = 4
	ring, ok = rings["Third"]
	fmt.Printf("ring: %d, ok: %v\n", ring, ok)
	delete(rings, "Third")
	ring, ok = rings["Third"]
	fmt.Printf("ring: %d, ok: %v\n", ring, ok)

	isPrime := make(map[int]bool)
	var prime bool
	isPrime[5] = true
	isPrime[1] = true
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)
	delete(isPrime, 5)
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)
	prime, ok = isPrime[1]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)

	counts := make(map[string]int)
	counts["Masha"]++
	counts["Vasya"]++
	counts["Sasha"]++
	counts["Masha"]++
	fmt.Println(counts)

	grades := map[string]float64{"Vasya": 54.2, "Alisa": 81.2, "Anna": 64.2}
	var names []string
	for name := range grades {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s имеет следующее количество балло %0.1f%%\n", name, grades[name])
	}
}
