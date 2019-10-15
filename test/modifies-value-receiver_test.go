package test

import (
	"testing"

	"github.com/calmisland/go-revive/rule"
)

func TestModifiesValRec(t *testing.T) {
	testRule(t, "modifies-value-receiver", &rule.ModifiesValRecRule{})
}
