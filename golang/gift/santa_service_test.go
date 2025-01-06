package gift

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluateRequestForNiceChild(t *testing.T) {
	niceChild := Child{"Alice Thomas", Nice, GiftRequest{"Bicycle", true, NiceToHave}}
	service := SantaService{}

	assert.True(t, service.EvaluateRequest(niceChild))
}

func TestEvaluateRequestForNaughtyChild(t *testing.T) {
	naughtyChild := Child{"Noa Thierry", Naughty, GiftRequest{"SomeToy", true, Dream}}
	service := SantaService{}

	assert.False(t, service.EvaluateRequest(naughtyChild))
}

func TestEvaluateRequestForInfeasibleGift(t *testing.T) {
	infeasibleGiftChild := Child{"Charlie Joie", Nice, GiftRequest{"AnotherToy", false, Dream}}
	service := SantaService{}

	assert.False(t, service.EvaluateRequest(infeasibleGiftChild))
}
