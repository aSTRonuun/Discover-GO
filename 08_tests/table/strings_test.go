package stringss

import (
	"strings"
	"testing"
)

const msgIndex = "%s (part: %s) - Indexs: expected (%d) <> result (%d)."

func TestIndex(t *testing.T) {
	t.Parallel()
	tests :=
		[]struct {
			text     string
			part     string
			expected int
		}{
			{"Cod3r is nice", "Cod3r", 0},
			{"", "", 0},
			{"Hey", "hey", -1},
			{"Vitor", "o", 3},
		}

	for _, test := range tests {
		t.Logf("Mass: %v", test)
		actual := strings.Index(test.text, test.part)

		if actual != test.expected {
			t.Errorf(msgIndex, test.text, test.part, test.expected, actual)
		}
	}
}
