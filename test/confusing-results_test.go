package test

import (
	"testing"

	"github.com/calmisland/go-revive/rule"
)

func TestConfusingResults(t *testing.T) {
	testRule(t, "confusing-results", &rule.ConfusingResultsRule{})
}
