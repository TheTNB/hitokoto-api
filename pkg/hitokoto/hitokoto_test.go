package hitokoto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {
	sentence, err := Random()

	assert.NoError(t, err)
	assert.NotEmpty(t, sentence)
}
