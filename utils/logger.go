package utils

import "testing"

func LogResult(t *testing.T, result int) {
	t.Logf("Result: %d", result)
}
