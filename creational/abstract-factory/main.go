package main

import "fmt"

/*
Say, you need to buy a sports kit, a set of two different products: a pair of shoes and a shirt. You would want to buy a full
sports kit of the same brand to match all the items.
If we try to turn this into code, the abstract factory will help us create sets of products so that they would always match each other.
*/

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.createShoes()
	nikeShirt := nikeFactory.createShirt()

	adidasShoe := adidasFactory.createShoes()
	adidasShirt := adidasFactory.createShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s IShoes) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
