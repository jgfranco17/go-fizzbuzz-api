package calculation

import (
	"fmt"
)

// Generate a fizzbuzz sequence up until the given number.
// Upperbound number must be a positive integer.
func GenerateFizzbuzzSequence(number int) ([]string, error) {
	if number <= 0 {
		return nil, fmt.Errorf("Number must be greater than zero")
	}
	var sequence []string
	for n := 1; n <= number; n++ {
		entry := ""
		if n%3 == 0 {
			entry += "FIZZ"
		}
		if n%5 == 0 {
			entry += "BUZZ"
		}

		if entry == "" {
			entry = fmt.Sprintf("%d", n)
		}
		sequence = append(sequence, entry)
	}
	return sequence, nil
}
