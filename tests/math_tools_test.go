package tests

import (
	"go-practice/math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMath(t *testing.T) {
	assert.Equal(t, math.Circle{Radius: 5}.Area(), 78.5, "they should be equal")
}
