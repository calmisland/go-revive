package test

import (
	"testing"

	"github.com/calmisland/go-revive/rule"
)

// Atomic rule.
func TestAtomic(t *testing.T) {
	testRule(t, "atomic", &rule.AtomicRule{})
}
