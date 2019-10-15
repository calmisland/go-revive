package test

import (
	"testing"

	"github.com/calmisland/go-revive/rule"
)

func TestUnreachableCode(t *testing.T) {
	testRule(t, "unreachable-code", &rule.UnreachableCodeRule{})
}
