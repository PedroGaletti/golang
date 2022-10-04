package main

import "fmt"

// var MyMap map[keyType]valType

func MapsDescription() {
	fmt.Println("\nMapsDescription func: ")
	heroes := make(map[string]string)
	villians := make(map[string]string)

	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["Hulk"] = "Bruce"
	villians["Wanda"] = "Vision"

	pets := map[int]string{1: "Dog", 2: "Cat"}
	fmt.Printf("Batman is: %v\n", heroes["Batman"]) // Batman is: Bruce Wayne
	fmt.Printf("Index 1 is: %v\n", pets[1])         // Index 1 is: Dog

	fmt.Println("Chip", pets[3]) //

	_, ok := pets[3]
	fmt.Println("Is there a 3rd pet: ", ok) // Is there a 3rd pet:  false

	for k, val := range heroes {
		fmt.Printf("%s is %s\n", k, val) // Batman is Bruce Wayne ......
	}

	fmt.Println("heroes", heroes) // heroes map[Batman:Bruce Wayne Hulk:Bruce Superman:Clark Kent]
	delete(heroes, "Batman")
	fmt.Println("heroes", heroes) // heroes map[Hulk:Bruce Superman:Clark Kent]
}

func main() {
	MapsDescription()
}
