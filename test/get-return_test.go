package test

import (
	"testing"

	"github.com/calmisland/go-revive/rule"
)

func TestGetReturn(t *testing.T) {
	testRule(t, "get-return", &rule.GetReturnRule{})
}
