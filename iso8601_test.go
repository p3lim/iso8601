package iso8601

import (
	"fmt"
	"time"
	"testing"
)

func TestISO8601Format(t *testing.T) {
	var d time.Duration = 54321 * time.Hour

	if Format(d) != "P6Y2M1W4D" {
		t.Fatalf("String format for %s does not equal expected %s", d, )
	}
}
