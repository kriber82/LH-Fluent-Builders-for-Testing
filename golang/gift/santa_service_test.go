package gift

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluateRequest(t *testing.T) {
	service := SantaService{}

	niceChild := Child{"Alice Thomas", Nice, GiftRequest{"Bicycle", true, NiceToHave}}
	assert.True(t, service.EvaluateRequest(niceChild))

	naughtyChild := Child{"Noa Thierry", Naughty, GiftRequest{"SomeToy", true, Dream}}
	assert.False(t, service.EvaluateRequest(naughtyChild))

	infeasibleGiftChild := Child{"Charlie Joie", Nice, GiftRequest{"AnotherToy", false, Dream}}
	assert.False(t, service.EvaluateRequest(infeasibleGiftChild))
}
