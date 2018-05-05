package main

func main() {
    cards := newDeck()
    // cards.print_deck()
    // hand, remainingCards := deal(cards, 5)

    // hand.print_deck()
    // remainingCards.print_deck()

    // fmt.Println(cards.toString())

    // cards.saveToFile("my_saved_deck")

    // Note: the variable cards remains unchanged.

    // deck := newDeckFromFile("my_saved_deck1")
    // deck.print_deck()

    cards.shuffle()
    cards.print_deck()
}

// deck -> []string -> string -> []byte
// the []byte is the type required by WriteToFile