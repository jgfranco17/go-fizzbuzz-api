package models

type Response struct {
	Limit    int      `json:"limit"`
	Fizz     int      `json:"fizz"`
	Buzz     int      `json:"buzz"`
	FizzBuzz int      `json:"fizzbuzz"`
	Digits   int      `json:"digits"`
	Sequence []string `json:"sequence"`
}
