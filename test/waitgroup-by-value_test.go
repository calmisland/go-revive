package test

import (
	"testing"

	"github.com/calmisland/go-revive/rule"
)

func TestWaitGroupByValue(t *testing.T) {
	testRule(t, "waitgroup-by-value", &rule.WaitGroupByValueRule{})
}
