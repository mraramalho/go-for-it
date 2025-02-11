package deck

import "testing"


func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	if len(deck.Cards) != 52 {
		t.Error("Expected 52 cards, got", len(deck.Cards))
	}
}

func TestDeckToString(t *testing.T) {
	deck := NewDeck()
	// Simular um deck com apenas duas cartas
	deck.CardCount = 2 // Atualiza o contador de cartas no deck
	deck.Cards[0] = Card{Value: "Ace", Suit: "Spades"}
	deck.Cards[1] = Card{Value: "Two", Suit: "Spades"}

	expected := "Ace of Spades,Two of Spades"
	result := deck.DeckToString()

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestAddCardToFullDeck(t *testing.T){
	deck := NewDeck()
	card := Card{Value: "Ace", Suit: "Spades"}
	if err := deck.AddCard(card); err == nil{
		t.Error("Expected error, got nil")
	}
}

func TestAddCardToEmptyDeck(t *testing.T){
	deck := Deck{
		Cards: [52]Card{},
		CardCount: 0,
	}
	card := Card{Value: "Ace", Suit: "Spades"}
	if err := deck.AddCard(card); err != nil{
		t.Error("Expected nil, got", err)
	}
}