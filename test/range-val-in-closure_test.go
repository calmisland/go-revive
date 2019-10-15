package test

import (
	"testing"

	"github.com/calmisland/go-revive/lint"
	"github.com/calmisland/go-revive/rule"
)

func TestRangeValInClosure(t *testing.T) {
	testRule(t, "range-val-in-closure", &rule.RangeValInClosureRule{}, &lint.RuleConfig{})
}
