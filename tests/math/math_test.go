package math

import "testing"

const defaultError = "Expected %v, got %v"

/*
	All Test methods must start with the word Test
	and must be exported (start with a capital letter)
*/
func TestAverage(t *testing.T) {
	expected := 7.28
	value := Average(7.2, 9.9, 6.1, 5.9)

	if value != expected {
		t.Errorf(defaultError, expected, value)
	}
}
