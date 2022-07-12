package glx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInf(t *testing.T) {
	assert.True(t, IsPositiveInf(InfPositive()))
	assert.False(t, IsPositiveInf(InfNegative()))
	assert.False(t, IsPositiveInf(Max))
	assert.False(t, IsPositiveInf(Max+1))

	assert.True(t, IsNegativeInf(InfNegative()))
	assert.False(t, IsNegativeInf(InfPositive()))
	assert.False(t, IsNegativeInf(Min))
	assert.False(t, IsNegativeInf(Min-1))
}
