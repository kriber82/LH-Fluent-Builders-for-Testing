package gift

type GiftRequest struct {
	giftName   string
	isFeasible bool
	priority   Priority
}

func NewGiftRequest(giftName string, isFeasible bool, priority Priority) *GiftRequest {
	return &GiftRequest{
		giftName:   giftName,
		isFeasible: isFeasible,
		priority:   priority,
	}
}

func (g *GiftRequest) GiftName() string {
	return g.giftName
}

func (g *GiftRequest) IsFeasible() bool {
	return g.isFeasible
}

func (g *GiftRequest) Priority() Priority {
	return g.priority
}
