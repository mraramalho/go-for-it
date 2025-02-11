package deck

import "testing"

func TestNewDeck(t *testing.T) {
	d := NewDeck()
	if len(d.Cards) != 52 {
		t.Error("Expected 52 cards, got", len(d.Cards))
	}
	if d.CardCount != 52 {
		t.Error("Expected 52 cards in CardCount, got", d.CardCount)
	}
}

func TestDeckToString(t *testing.T) {
	d := NewDeck()
	// Simular um deck com apenas duas cartas
	d.CardCount = 2 // Atualiza o contador de cartas no deck
	d.Cards[0] = Card{Value: "Ace", Suit: "Spades"}
	d.Cards[1] = Card{Value: "Two", Suit: "Spades"}

	expected := "Ace of Spades,Two of Spades"
	result := d.DeckToString()

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestAddCardToFullDeck(t *testing.T) {
	d := NewDeck()
	card := Card{Value: "Ace", Suit: "Spades"}
	if err := d.AddCard(card); err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestAddCardToEmptyDeck(t *testing.T) {
	d := NewEmptyDeck()
	card := Card{Value: "Ace", Suit: "Spades"}
	if err := d.AddCard(card); err != nil {
		t.Error("Expected nil, got", err)
	}
	if d.CardCount != 1 {
		t.Error("Expected 1 card, got", d.CardCount)
	}
}

func TestReadDeckFromFile(t *testing.T) {
	d := NewDeck()
	temp_dir := t.TempDir()
	d.SaveDeckToFile(temp_dir + "temp_file.txt")
	d, err := ReadDeckFromFile(temp_dir + "temp_file.txt")
	if err != nil {
		t.Error("Expected nil, got", err)
	}
	if d.CardCount != 52 {
		t.Error("Expected 52 cards, got", d.CardCount)
	}
}

func TestReadDeckFromFileError(t *testing.T) {
	d := Deck{
		Cards:     [52]Card{},
		CardCount: 0,
	}
	temp_dir := t.TempDir()
	d.SaveDeckToFile(temp_dir + "temp_file.txt")
	d, err := ReadDeckFromFile(temp_dir + "temp_file.txt")
	if err == nil {
		t.Error("no cards in file: "+temp_dir+"temp_file.txt, got: ", err)
	}

}

func TestDeckShuffle(t *testing.T) {
	d1, d2 := NewDeck(), NewDeck()
	d1.Shuffle()
	d2.Shuffle()
	if d1.DeckToString() == d2.DeckToString() {
		t.Error("Expectted different shuffled decks, got the same")
	}
}
