package main

func main(){
	cards := deck{newCart("Tomas"), newCart("Martin")}
	cards = append(cards, newCart("John"))
	cards.print()
}

func newCart(val string) string {
	return "Hi " + val
}