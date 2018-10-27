package iso8601

import (
	"testing"
	"time"
)

func TestISO8601Format(t *testing.T) {
	tests := map[time.Duration]string{
		54321 * time.Hour:   "P6Y2M1W4D",
		54321 * time.Minute: "P1M1WT6H51M",
		54321 * time.Second: "PT15H5M21S",
		0 * time.Second:     "PT0S",
	}

	for input, expected := range tests {
		output := Format(input)
		if output != expected {
			t.Fatalf("String format for %s does not equal expected %s, got %s", input, expected, output)
		}
	}
}
