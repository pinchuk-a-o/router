package router

import (
	"strconv"
	"testing"
)

func TestRuleInteger_Handler(t *testing.T) {
	provider := map[string]bool{"321": true, "A321": false, "-321": true, "321.0": false}

	r := RuleInteger{}

	for i, check := range provider {
		if res := r.Check(i); res != check {
			t.Error("(int)" + i + " is not " + strconv.FormatBool(check))
		}
	}
}

func TestRuleLetter_Handler(t *testing.T) {
	provider := map[string]bool{"321": false, "A321": false, "asddsa": true, "asd dsa": false}

	r := RuleLetter{}

	for i, check := range provider {
		if res := r.Check(i); res != check {
			t.Error("(string)" + i + " is not " + strconv.FormatBool(check))
		}
	}
}
