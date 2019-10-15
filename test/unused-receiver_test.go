package test

import (
	"testing"

	"github.com/calmisland/go-revive/rule"
)

func TestUnusedReceiver(t *testing.T) {
	testRule(t, "unused-receiver", &rule.UnusedReceiverRule{})
}
