package gift

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluateRequestForNiceChild(t *testing.T) {
	niceChild := NewChild("Alice", "Thomas", 9, Nice,
		*NewGiftRequest("Bicycle", true, NiceToHave))
	service := SantaService{}

	assert.True(t, service.EvaluateRequest(*niceChild))
}

func TestEvaluateRequestForNaughtyChild(t *testing.T) {
	naughtyChild := NewChild("Noa", "Thierry", 6, Naughty,
		*NewGiftRequest("SomeToy", true, Dream))
	service := SantaService{}

	assert.False(t, service.EvaluateRequest(*naughtyChild))
}

func TestEvaluateRequestForInfeasibleGift(t *testing.T) {
	infeasibleGiftChild := NewChild("Charlie", "Joie", 3, Nice,
		*NewGiftRequest("AnotherToy", false, Dream))
	service := SantaService{}

	assert.False(t, service.EvaluateRequest(*infeasibleGiftChild))
}
