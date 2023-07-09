package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"alice":   32,
		"charlie": 34,
	}

	newAges := make(map[string]int)
	newAges["alice"] = 32
	newAges["charlie"] = 34

	fmt.Printf("ages: %v\n", ages)
	fmt.Printf("newAges[\"alice\"]: %v\n", newAges["alice"])

	delete(ages, "charlie")
	fmt.Printf("ages: %v\n", ages)

	fmt.Println(ages["vini"])
	fmt.Printf("ages: %v\n", ages)

	ages["bob"]++
	fmt.Printf("ages: %v\n", ages)
	ages["vini"]++
	ages["baggio"]++
	// test := &ages["bob"] // compile error: cannot take address of map element

	for key, value := range ages {
		fmt.Printf("Key: %s\tValues: %d\n", key, value)
	}
	fmt.Println("")
	var names []string
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("Key: %s\tValues: %d\n", name, ages[name])
	}

	fmt.Printf("len(ages): %v\n", len(ages))

	var age map[string]int
	fmt.Println(age == nil)    // "true"
	fmt.Println(len(age) == 0) // "true"
	// age["vini"] = 0 panic: assignment to entry in nil map

	// age_, ok := ages["test"]
	// if !ok {
	// 	fmt.Printf("key: %s is not on the map but age is zero value for the type %d\n", "test", age_)
	// }

	if age_, ok := ages["test"]; !ok {
		fmt.Printf("key: %s is not on the map but age is zero value for the type %d\n", "test", age_)
	}
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}
