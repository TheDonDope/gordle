package storage

import (
	"fmt"
	"testing"
)

func TestNewWotdSucceeds(t *testing.T) {
	want := 5
	got := len(NewWotd())

	if got != want {
		t.Error(fmt.Printf("Should have matched, got: %v, want: %v", got, want))
	}
}
