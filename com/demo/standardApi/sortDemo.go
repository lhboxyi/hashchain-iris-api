package standardApi

import (
	"fmt"
	"sort"
)

func SortDemo() {
	ints := []int{4, 3, 1, 6, 0}
	sort.Ints(ints)
	fmt.Println(ints)

	floats := []float64{4.1, 3.5, 1.1, 6.2, 0.0}
	sort.Float64s(floats)
	fmt.Println(floats)
}

func StructSort() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println(people)
}
