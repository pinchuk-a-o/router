package router

import (
	"strconv"
	"unicode"
)

// Rule ...
type Rule interface {
	Check(variable string) bool
}

// Rules ...
type Rules struct {
	data map[string]Rule
}

func (r *Rules) compare(caseName, caseData string) bool {
	rule, ok := r.data[caseName]

	if !ok {
		return true
	}

	return rule.Check(caseData)
}

func (r *Rules) setRules(rules map[string]Rule) {
	r.data = rules
}

// RuleInteger ...
type RuleInteger struct {
}

// Check ...
func (r RuleInteger) Check(variable string) bool {
	_, err := strconv.Atoi(variable)

	return err == nil
}

// RuleLetter ...
type RuleLetter struct {
}

// Check ...
func (r RuleLetter) Check(variable string) bool {
	for _, v := range variable {
		if !unicode.IsLetter(v) {
			return false
		}
	}

	return true
}
