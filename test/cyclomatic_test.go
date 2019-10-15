package test

import (
	"testing"

	"github.com/calmisland/go-revive/lint"
	"github.com/calmisland/go-revive/rule"
)

func TestCyclomatic(t *testing.T) {
	testRule(t, "cyclomatic", &rule.CyclomaticRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(1)},
	})
	testRule(t, "cyclomatic-2", &rule.CyclomaticRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(3)},
	})
}
