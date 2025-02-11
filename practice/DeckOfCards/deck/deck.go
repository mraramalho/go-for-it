package deck

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Deck struct {
	Cards     [52]Card
	CardCount int // Contador de cartas atuais no baralho
}

// Imprime todas as cartas no baralho
func (d *Deck) Print() {
	for i := 0; i < d.CardCount; i++ {
		fmt.Printf("%s of %s\n", d.Cards[i].Value, d.Cards[i].Suit)
	}
}

// Adiciona uma carta ao baralho se ainda houver espaço
func (d *Deck) AddCard(card Card) error {
	if d.CardCount >= 52 {
		return errors.New("can't add card, because deck is full")
	}
	d.Cards[d.CardCount] = card
	d.CardCount++
	return nil
}

// Retorna o baralho como uma string (para salvar ou imprimir)
func (d *Deck) DeckToString() string {
	cards := []string{}
	for i := 0; i < d.CardCount; i++ {
		cards = append(cards, d.Cards[i].ToString())
	}
	return strings.Join(cards, ",")
}

// Remove `handsize` cartas do baralho e retorna um novo Deck com essas cartas
func (d *Deck) Deal(handsize int) (Deck, error) {
	if handsize > d.CardCount {
		return Deck{}, errors.New("not enough cards to deal")
	}
	hand := Deck{}
	for i := range handsize {
		hand.AddCard(d.Cards[i])
	}
	copy(d.Cards[:], d.Cards[handsize:d.CardCount]) // Remove as cartas do baralho original
	d.CardCount -= handsize
	return hand, nil
}

// Embaralha as cartas no baralho
func (d *Deck) Shuffle() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range d.Cards[:d.CardCount] {
		j := rand.Intn(d.CardCount)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

// Salva o baralho em um arquivo
func (d *Deck) SaveDeckToFile(filepath string) error {
	if err := os.WriteFile(filepath, []byte(d.DeckToString()), 0666); err != nil {
		return err
	}
	return nil
}

// Cria um novo baralho padrão de 52 cartas
func NewDeck() Deck {
	d := Deck{}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	for _, suit := range suits {
		for _, value := range values {
			d.AddCard(Card{Value: value, Suit: suit})
		}
	}
	return d
}

func NewEmptyDeck() Deck {
	return Deck{
		Cards:     [52]Card{},
		CardCount: 0,
	}
}

// Lê um baralho de um arquivo
func ReadDeckFromFile(filepath string) (Deck, error) {
	bs, err := os.ReadFile(filepath)
	if err != nil {
		return Deck{}, err
	}
	s := strings.Split(string(bs), ",")
	d := Deck{}
	for _, card := range s {
		cardParts := strings.Split(strings.TrimSpace(card), " of ")
		if len(cardParts) != 2 {
			return Deck{}, fmt.Errorf("invalid card format: %s", card)
		}
		d.AddCard(Card{
			Value: strings.TrimSpace(cardParts[0]),
			Suit:  strings.TrimSpace(cardParts[1]),
		})
	}
	return d, nil
}
