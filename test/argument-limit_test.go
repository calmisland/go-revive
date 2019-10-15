package test

import (
	"testing"

	"github.com/calmisland/go-revive/lint"
	"github.com/calmisland/go-revive/rule"
)

func TestArgumentLimit(t *testing.T) {
	testRule(t, "argument-limit", &rule.ArgumentsLimitRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(3)},
	})
}
