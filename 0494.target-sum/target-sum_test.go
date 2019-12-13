package _494_target_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindTargetSumWays(t *testing.T) {
	assert.Equal(t, findTargetSumWays([]int{1, 1, 1, 1, 1}, 3), 5)
}
