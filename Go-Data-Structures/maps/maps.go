package main

import (
	"fmt"
	"sort"
)

func main() {
	nationCapitalsExample()
}

func nationCapitalsExample() {

	// Note: In Go, maps have no contract to maintain the order of the keys
	var nationCapitals = make(map[string]string)
	nationCapitals["Afganistan"] = "Kabul"
	nationCapitals["Canada"] = "Ottawa"
	nationCapitals["Japan"] = "Tokyo"
	nationCapitals["Kenya"] = "Nairobi"
	nationCapitals["India"] = "New Delhi"
	nationCapitals["Mexico"] = "Mexico City"
	nationCapitals["South Korea"] = "Seoul"
	nationCapitals["United Kindom"] = "London"
	nationCapitals["USA"] = "Washington D.C."
	nationCapitals["Taiwan"] = "Taipei"

	fmt.Println("Print the map unsorted (random order): ")
	printNationCapitalsMap(nationCapitals)

	fmt.Println("Print themap sorted by key (nation name):")
	printSortedNationCapitalsMap(nationCapitals)
}

func printNationCapitalsMap(capitalsMap map[string]string) {
	for key, value := range capitalsMap {
		fmt.Println("The capital of", key, "is", value)
	}
}

// An example of sorting the keys in a map alphabetically
func printSortedNationCapitalsMap(capitalsMap map[string]string) {
	keys := make([]string, len(capitalsMap))
	for key := range capitalsMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, v := range keys {
		if v == "" {
			continue
		}
		fmt.Println("The capital of", v, "is", capitalsMap[v])
	}
}
