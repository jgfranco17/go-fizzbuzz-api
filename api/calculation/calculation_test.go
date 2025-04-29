package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzbuzzSequenceValid(t *testing.T) {
	testCases := []struct {
		description string
		number      int
		expected    []string
	}{
		{
			description: "with fizz only",
			number:      3,
			expected:    []string{"1", "2", "FIZZ"},
		},
		{
			description: "with fizz, buzz, and fizzbuzz",
			number:      15,
			expected:    []string{"1", "2", "FIZZ", "4", "BUZZ", "FIZZ", "7", "8", "FIZZ", "BUZZ", "11", "FIZZ", "13", "14", "FIZZBUZZ"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			sequence, err := GenerateFizzbuzzSequence(tc.number)
			assert.NoErrorf(t, err, "Unexpected error for sequence %s", tc.description)
			assert.Equal(t, tc.expected, sequence)
		})
	}
}

func TestTestFizzbuzzSequenceInvalidInput(t *testing.T) {
	sequence, err := GenerateFizzbuzzSequence(-1)
	assert.ErrorContains(t, err, "Number must be greater than zero")
	assert.Nil(t, sequence)
}

func FuzzFizzbuzzSequence(f *testing.F) {
	validInputs := []int{1, 10, 100, 1000}
	for _, num := range validInputs {
		f.Add(num)
	}
	f.Fuzz(func(t *testing.T, input int) {
		if input <= 0 {
			t.Skip("Skipping non-positive input")
		}
		sequence, err := GenerateFizzbuzzSequence(input)
		assert.NoErrorf(t, err, "Unexpected error for sequence input %d", input)
		assert.NotEmpty(t, sequence, "Wanted sequence but got none")
	})
}

func BenchmarkFizzbuzzSequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GenerateFizzbuzzSequence(i)
		if err != nil {
			b.Fail()
		}
	}
}
