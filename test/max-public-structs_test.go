package test

import (
	"testing"

	"github.com/calmisland/go-revive/lint"
	"github.com/calmisland/go-revive/rule"
)

func TestMaxPublicStructs(t *testing.T) {
	testRule(t, "max-public-structs", &rule.MaxPublicStructsRule{}, &lint.RuleConfig{
		Arguments: []interface{}{int64(1)},
	})
}
