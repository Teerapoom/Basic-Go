package go_map

import "fmt"

func map_basic() {
	var numbers = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("", numbers)

	// Make Map to Dynamic
	var colors = make(map[string]string)
	colors["red"] = "#0f00"
	colors["green"] = "#0f0"
	colors["blue"] = "#00f"
	fmt.Println("", colors)        // map[ red = #0f00   green = #0f0 red =#0f00]
	fmt.Println("", colors["red"]) // #0f00

	// Mutiple Map 2 มิติ  nil now
	var courses = make(map[string]map[string]int)

	courses["Android"] = make(map[string]int) //กำหนดค่า
	courses["Android"]["price"] = 200
	courses["Android"]["code"] = 200
	fmt.Println(courses)
	// map[Android:map[price:200 code:200 ]]

}
