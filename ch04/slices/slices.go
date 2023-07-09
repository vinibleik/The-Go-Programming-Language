package main

import "fmt"

func main() {
	months := [...]string{"", "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Printf("months: %v\nlen(Q2): %v\tcap(Q2): %v\n\n", months, len(months), cap(months))

	fmt.Printf("Q2: %v\nlen(Q2): %v\tcap(Q2): %v\n\n", Q2, len(Q2), cap(Q2))
	fmt.Printf("summer: %v\nlen(summer): %v\tcap(summer): %v\n\n", summer, len(summer), cap(summer))

	// fmt.Println(summer[:20]) // panic: out of range

	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Printf("endlessSummer: %v\nlen(endlessSummer): %v\tcap(endlessSummer): %v\n\n", endlessSummer, len(endlessSummer), cap(endlessSummer))
	newSummer := summer[:]
	fmt.Printf("newSummer: %v\nlen(newSummer): %v\tcap(newSummer): %v\n\n", newSummer, len(newSummer), cap(newSummer))
}
