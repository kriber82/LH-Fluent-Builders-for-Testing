package gift

type SantaService struct{}

func (s *SantaService) EvaluateRequest(child Child) bool {
	return child.Behavior == Nice && child.GiftRequest.IsFeasible
}
