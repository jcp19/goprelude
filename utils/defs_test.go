package utils_test

import (
	"testing"

	"github.com/jcp19/prelude/utils"
	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	// create two slices already at capacity
	x1 := make([]int, 3, 3)
	x2 := make([]int, 3, 3)
	// 'append'ing to a slice at capacity causes
	// a new slice to be allocated; x1 still contains
	// the original here.
	_ = append(x1, 1)
	assert.Equal(t, len(x1), 3)
	// 'utils.Append'ing to a slice requires a pointer
	// to the slice and updates the reference if need be.
	utils.Append(&x2, 1)
	assert.Equal(t, len(x2), 4)
}
