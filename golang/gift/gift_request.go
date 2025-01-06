package gift

type GiftRequest struct {
	GiftName   string
	IsFeasible bool
	Priority   Priority
}

func (g *GiftRequest) GetGiftName() string {
	return g.GiftName
}

func (g *GiftRequest) IsGiftFeasible() bool {
	return g.IsFeasible
}

func (g *GiftRequest) GetPriority() Priority {
	return g.Priority
}
