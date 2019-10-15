package test

import (
	"testing"

	"github.com/calmisland/go-revive/rule"
)

// Tests RedefinesBuiltinID rule.
func TestRedefinesBuiltinID(t *testing.T) {
	testRule(t, "redefines-builtin-id", &rule.RedefinesBuiltinIDRule{})
}
