package test

import (
	"testing"

	"github.com/calmisland/go-revive/rule"
)

func TestImportShadowing(t *testing.T) {
	testRule(t, "import-shadowing", &rule.ImportShadowingRule{})
}
