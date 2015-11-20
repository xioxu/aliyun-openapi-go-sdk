package core

import "testing"

func expectPanic(t *testing.T, f func()) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("Expect panic but no error were got")
		}
	}()
	f()
}
