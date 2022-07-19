package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRuleString(t *testing.T) {
	// Valid rule
	rule, err := parseRuleString("cpu.all >= 50 : CPU should not go over 50%")
	assert.NoError(t, err)
	assert.Equal(t, rule.name, "CPU should not go over 50%")
	assert.Equal(t, rule.Operator, ">=")
	assert.Equal(t, rule.Property, "cpu.all")
	assert.EqualValues(t, rule.Value, 50)

	// Missing name resolves to rule
	rule, err = parseRuleString("cpu.all >= 50")
	assert.NoError(t, err)
	assert.Equal(t, rule.name, "cpu.all >= 50")

	// Missing elements
	rule, err = parseRuleString("cpu.all 50")
	assert.ErrorContains(t, err, "does not adhere to the format")

	// Invalid operator
	rule, err = parseRuleString("cpu.all => 50")
	assert.ErrorContains(t, err, "does not match allowed operators")

	// Non numeric val
	rule, err = parseRuleString("cpu.all > cat")
	assert.ErrorContains(t, err, "contained a non-numeric value")
}
