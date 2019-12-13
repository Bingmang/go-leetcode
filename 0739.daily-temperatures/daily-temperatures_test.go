package _739_daily_temperatures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	assert.Equal(t, dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}), []int{1, 1, 4, 2, 1, 1, 0, 0})
}
