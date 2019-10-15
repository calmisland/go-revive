package test

import (
	"testing"

	"github.com/calmisland/go-revive/lint"
	"github.com/calmisland/go-revive/rule"
)

func TestFunctionResultsLimit(t *testing.T) {
	testRule(t, "function-result-limit", &rule.FunctionResultsLimitRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(3)},
	})
}
