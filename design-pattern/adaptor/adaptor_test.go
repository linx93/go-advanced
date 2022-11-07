package adaptor

import (
	"testing"
)

func TestNewAdaptee(t *testing.T) {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)

	res := target.Request() // adaptee method

	t.Fatalf("res: %s", res)

}
