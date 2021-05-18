package dopo

import (
	"testing"
	"time"
)

func TestAwait(t *testing.T) {
	var ctrl int
	dopo := Exec(func() interface{} {
		c := 123
		time.Sleep(100 * time.Millisecond)
		return c
	})

	ctrl = 321

	result := dopo.await()
	if result != 123 {
		t.Errorf("Expected result to be %v, got %v", 123, result)
	}

	if ctrl != 321 {
		t.Errorf("Expected ctrl to be %v, got %v", 321, ctrl)
	}

}
