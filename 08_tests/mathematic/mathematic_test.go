package mathematic

import "testing"

const standardError = "The expected value was %v, but the result found was %v"

func TestAvarage(t *testing.T) {
	t.Parallel()
	expectedValue := 7.28
	value := Avarage(7.2, 9.9, 6.1, 5.9)

	if value != expectedValue {
		t.Errorf(standardError, expectedValue, value)
	}
}
