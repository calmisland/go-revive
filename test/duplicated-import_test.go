package test

import (
	"testing"

	"github.com/calmisland/go-revive/rule"
)

func TestDuplicatedImports(t *testing.T) {
	testRule(t, "duplicated-imports", &rule.DuplicatedImportsRule{})
}
