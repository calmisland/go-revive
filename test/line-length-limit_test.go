package test

import (
	"testing"

	"github.com/calmisland/go-revive/lint"
	"github.com/calmisland/go-revive/rule"
)

func TestLineLengthLimit(t *testing.T) {
	testRule(t, "line-length-limit", &rule.LineLengthLimitRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(100)},
	})
}
