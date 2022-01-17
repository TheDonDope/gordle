package storage

import (
	"fmt"
	"testing"
)

func TestNewWotdSucceeds(t *testing.T) {
	d := NewDictionary(5)
	want := 5
	got := len(d.NewWotd())

	if got != want {
		t.Error(fmt.Printf("Should have matched, got: %v, want: %v", got, want))
	}
}
