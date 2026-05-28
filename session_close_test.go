//go:build windows
// +build windows

package etw

import (
	"errors"
	"testing"
)

func TestCloseSessionStopsSessionWhenUnsubscribeFails(t *testing.T) {
	unsubscribeErr := errors.New("disable provider failed")
	stopped := false

	err := closeSession(
		func() error { return unsubscribeErr },
		func() error {
			stopped = true
			return nil
		},
	)

	if !stopped {
		t.Fatal("stop was not called after unsubscribe failed")
	}
	if !errors.Is(err, unsubscribeErr) {
		t.Fatalf("expected unsubscribe error, got %v", err)
	}
}
