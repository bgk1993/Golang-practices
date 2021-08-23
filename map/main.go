package main

import "fmt"

func main() {

	mapValues := map[int]string{
		1: "Bhargava Kulkarni",
		2: "Bhavani Kulkarni",
		3: "Gururaj Kulkarni",
		4: "Bheemabai Kulkarni",
	}

	delete(mapValues, 1)

	printMap(mapValues)
}

func printMap(mapValues map[int]string) {
	for key, value := range mapValues {
		fmt.Println("Key = ", key, ", Value = ", value)
	}
}
