package main

import "fmt"

func main(){
	//cards := deck{newCart("Tomas"), newCart("Martin")}
	//cards = append(cards, newCart("John"))
	cards := newDeck()
	fmt.Println(cards.toString())
	fmt.Println([]byte(cards.toString()))
	cards.saveToFile("test.txt")
	// hand, remainingCards := deal(cards, 2)
	
	// hand.print()
	// remainingCards.print()

	//cards.print()

	greeting := "Hi there!"
	fmt.Println([]byte(greeting))
}


// func newCart(val string) string {
// 	return "Hi " + val
// }