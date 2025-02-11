package deck

import (
	"testing"
)

func TestCardPrint(t *testing.T) {
	card := Card{Value: "Ace", Suit: "Spades"}
	card.ToString()
	if card.ToString() != "Ace of Spades" {
		t.Error("Expected Ace of Spades, got", card.ToString())
	}
}
