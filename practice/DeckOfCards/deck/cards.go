package deck

type Card struct {
	Value string
	Suit  string
}

func (c *Card) ToString() string {
	return c.Value + " of " + c.Suit
}
