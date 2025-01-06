package gift

type Child struct {
	firstName   string
	lastName    string
	age         int
	behavior    Behavior
	giftRequest GiftRequest
}

func NewChild(firstName, lastName string, age int, behavior Behavior, giftRequest GiftRequest) *Child {
	return &Child{
		firstName:   firstName,
		lastName:    lastName,
		age:         age,
		behavior:    behavior,
		giftRequest: giftRequest,
	}
}

func (c *Child) FirstName() string {
	return c.firstName
}

func (c *Child) LastName() string {
	return c.lastName
}

func (c *Child) Age() int {
	return c.age
}

func (c *Child) Behavior() Behavior {
	return c.behavior
}

func (c *Child) GiftRequest() GiftRequest {
	return c.giftRequest
}
