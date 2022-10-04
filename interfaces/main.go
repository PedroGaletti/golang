package main

import "fmt"

type Animal interface {
	AngrySound()
	HappySound()
}

type Cat string

func (c Cat) Attack() {
	fmt.Println("Cat attacks it's Prey")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	fmt.Println("Cat say hisss")
}

func (c Cat) HappySound() {
	fmt.Println("Cat say hisss")
}

func main() {
	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()

	var bob Cat = kitty.(Cat)
	bob.Attack()
	fmt.Println("Cats name: ", bob.Name()) // Cats name: Kitty
}
