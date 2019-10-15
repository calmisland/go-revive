package test

import (
	"testing"

	"github.com/calmisland/go-revive/rule"
)

// TestUnnecessaryStmt rule.
func TestUnnecessaryStmt(t *testing.T) {
	testRule(t, "unnecessary-stmt", &rule.UnnecessaryStmtRule{})
}
