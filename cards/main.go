package main

func main(){
	cards := deck{newCart("tomas"), newCart("Martin")}
	cards = append(cards, newCart("jozka"))
	deck.print(cards)
}

func newCart(val string) string {
	return "Hi " + val
}