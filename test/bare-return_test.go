package test

import (
	"testing"

	"github.com/calmisland/go-revive/rule"
)

func TestBareReturn(t *testing.T) {
	testRule(t, "bare-return", &rule.BareReturnRule{})
}
