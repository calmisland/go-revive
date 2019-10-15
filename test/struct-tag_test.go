package test

import (
	"testing"

	"github.com/calmisland/go-revive/rule"
)

// TestStructTag tests struct-tag rule
func TestStructTag(t *testing.T) {
	testRule(t, "struct-tag", &rule.StructTagRule{})
}
