package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	//var names CollectionOfNames

	//names = newDeck()
	//names.printNames()

	// fmt.Println()

	// first, second := deal(names, 8)

	// first.printNames()
	// fmt.Println()
	// second.printNames()

	//val := names.toString()

	// Printing appended string
	// fmt.Println(val)

	// printing ASCII bytes array
	// fmt.Println(([]byte)val)

	// names.saveToFile("my_file")
	// names = newCollectionFromFile("my_file")
	// names.printNames()

	values := newInts()
	for value := range values {
		if value%2 == 0 {
			fmt.Println(value, " is Even")
		} else {
			fmt.Println(value, " is Odd")
		}
	}
}

type CollectionOfNames []string
type CollectionOfInts []int

func newDeck() CollectionOfNames {
	first := CollectionOfNames{"first", "second", "third", "fourth"}
	second := CollectionOfNames{"fifth", "sixth", "seventh", "eighth"}

	third := CollectionOfNames{}

	for _, f := range first {
		for _, s := range second {
			third = append(third, f+" "+s)
		}
	}

	return third
}

func newInts() CollectionOfInts {
	values := CollectionOfInts{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	return values
}

func (names CollectionOfNames) printNames() {
	for _, name := range names {
		fmt.Println(name)
	}
}

func deal(names CollectionOfNames, number int) (CollectionOfNames, CollectionOfNames) {
	return names[:number], names[number:]
}

/*
 * Function to convert slice of strings to single string
 * resulting string will be returned from this function
 */
func (names CollectionOfNames) toString() string {
	return strings.Join(names, ",")
}

func (names CollectionOfNames) saveToFile(filename string) {
	str := names.toString()
	ioutil.WriteFile(filename, []byte(str), 0666)
}

func newCollectionFromFile(filename string) CollectionOfNames {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	r := strings.Split(string(bytes), ",")
	return CollectionOfNames(r)
}
