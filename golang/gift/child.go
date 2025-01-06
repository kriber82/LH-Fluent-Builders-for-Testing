package gift

type Child struct {
	Name       string
	Behavior   Behavior
	GiftRequest GiftRequest
}

func NewChild(firstName, lastName string, behavior Behavior, giftRequest GiftRequest) *Child {
	return &Child{
		Name:       firstName + " " + lastName,
		Behavior:   behavior,
		GiftRequest: giftRequest,
	}
}

func (c *Child) GetName() string {
	return c.Name
}

func (c *Child) GetBehavior() Behavior {
	return c.Behavior
}

func (c *Child) GetGiftRequest() GiftRequest {
	return c.GiftRequest
}