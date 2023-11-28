package util

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReverseIntNegative(t *testing.T) {
	val, err := ReverseInt(-629)

	require.NoError(t, err)
	assert.Equal(t, -926, val)
}

func TestReverseIntWrongInput(t *testing.T) {
	_, err := ReverseInt("612")
	assert.Error(t, err)
}

func TestReverseStringZero(t *testing.T) {
	val, err := ReverseInt(0)

	require.NoError(t, err)
	require.Equal(t, 0, val)
}

func TestContainsDuplicate(t *testing.T) {
	req := require.New(t)

	cases := map[string]struct {
		value []int
		want  bool
	}{
		"without duplicates": {value: []int{1, 2, 3, 4, 5}, want: false},
		"with duplicates":    {value: []int{1, 2, 2, 4, 8}, want: true},
		"empty":              {value: []int{}, want: false},
	}

	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			res := ContainsDuplicate(testCase.value)
			req.Equal(testCase.want, res)
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	req := require.New(t)
	testSuccess := func(value int, want bool) func(t *testing.T) {
		return func(t *testing.T) {
			res := IsPalindrome(value)
			req.Equal(want, res)
		}
	}

	t.Run("Zero value", testSuccess(0, true))
	t.Run("Positive palindrome", testSuccess(636, true))
	t.Run("Positive not palindrome", testSuccess(120, false))
	t.Run("Negative value", testSuccess(-5, false))
}
