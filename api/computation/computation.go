package computation

import (
	"fmt"

	"api/models"
)

func fizzbuzz(number int) string {
	result := ""
	if number%3 == 0 {
		result += "fizz"
	}
	if number%5 == 0 {
		result += "buzz"
	}
	if result != "" {
		return result
	}
	return fmt.Sprintf("%d", number)
}

func CalculateFizzbuzz(limit int) models.Response {
	var sequence []string
	fizzbuzzResponse := models.Response{
		Limit:    limit,
		Fizz:     0,
		Buzz:     0,
		FizzBuzz: 0,
	}
	for i := 1; i <= limit; i++ {
		val := fizzbuzz(i)
		if val == "fizz" {
			fizzbuzzResponse.Fizz += 1
		} else if val == "buzz" {
			fizzbuzzResponse.Buzz += 1
		} else if val == "fizzbuzz" {
			fizzbuzzResponse.FizzBuzz += 1
		} else {
			fizzbuzzResponse.Digits += 1
		}
		sequence = append(sequence, val)
	}
	fizzbuzzResponse.Sequence = sequence
	return fizzbuzzResponse
}
