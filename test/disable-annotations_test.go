package test

import (
	"testing"

	"github.com/calmisland/go-revive/lint"
	"github.com/calmisland/go-revive/rule"
)

func TestDisabledAnnotations(t *testing.T) {
	testRule(t, "disable-annotations", &rule.ExportedRule{}, &lint.RuleConfig{})
}

func TestModifiedAnnotations(t *testing.T) {
	testRule(t, "disable-annotations2", &rule.VarNamingRule{}, &lint.RuleConfig{})
}
