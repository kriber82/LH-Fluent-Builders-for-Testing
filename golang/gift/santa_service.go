package gift

type SantaService struct{}

func (s *SantaService) EvaluateRequest(child Child) bool {
	return child.behavior == Nice && child.giftRequest.isFeasible
}
