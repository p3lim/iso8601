package iso8601

import (
	"time"
	"testing"
)

func TestISO8601Format(t *testing.T) {
	var d time.Duration = 54321 * time.Hour

	expected := "P6Y2M1W4D"
	if Format(d) != expected {
		t.Fatalf("String format for %s does not equal expected %s", d, expected)
	}
}
