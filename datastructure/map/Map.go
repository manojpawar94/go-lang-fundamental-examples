package main

import "fmt"

func main() {
	// declaring the map using var keyword
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["India"] = "New Delhi"
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	// print the map
	fmt.Println("countryCapitalMap:", countryCapitalMap)

	// declaring the map without var keyword
	colorCodeMap := make(map[string]string)
	colorCodeMap["Black"] = "#000000"
	colorCodeMap["White"] = "#ffffff"
	fmt.Println("colorCodeMap:", colorCodeMap)

	// initializing map while declaring
	digitStrMap := map[int]string{1: "One", 2: "Two", 3: "Three", 4: "Four", 5: "Five", 6: "Six", 7: "Seven", 8: "Eight", 9: "Nine"}
	fmt.Println("digitStrMap:", digitStrMap)

	// adding new element to existing map
	digitStrMap[0] = "Zero"

	// element of map after adding the new element
	fmt.Println("digitStrMap:", digitStrMap)

	// iterating the map
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	// check if key is present in the map
	colorCode, isPresent := colorCodeMap["Black"]
	if isPresent {
		fmt.Println("The color code for Black color is", colorCode)
	} else {
		fmt.Println("The color code is not present for  Black color in Map")
	}

	// check if key is present in the map
	colorCode, isPresent = colorCodeMap["Blue"]
	if isPresent {
		fmt.Println("The color code for Blue  color is", colorCode)
	} else {
		fmt.Println("The color code is not present for Blue colorin Map")
	}

	// delete function on the map
	// iterating the map before deleting france from map
	fmt.Println("countryCapitalMap before deleting france from map...")
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	delete(countryCapitalMap, "France")

	// iterating the map after deleting france from map
	fmt.Println("countryCapitalMap after deleting france from map...")
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
}