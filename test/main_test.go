package test

import (
	"testing"

	"github.com/carrey1994/pipeline_project/api/math"
	"github.com/go-playground/assert/v2"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, math.Add(1, 2), 3)
}
