package test

import (
	"testing"

	"github.com/calmisland/go-revive/rule"
)

// TestModifiesParam rule.
func TestModifiesParam(t *testing.T) {
	testRule(t, "modifies-param", &rule.ModifiesParamRule{})
}
