package main

func main(){
	cards := deck{newCart("Tomas"), newCart("Martin")}
	cards = append(cards, newCart("John"))
	deck.print(cards)
}

func newCart(val string) string {
	return "Hi " + val
}