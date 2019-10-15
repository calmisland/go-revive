package test

import (
	"testing"

	"github.com/calmisland/go-revive/rule"
)

func TestDeepExit(t *testing.T) {
	testRule(t, "deep-exit", &rule.DeepExitRule{})
}
