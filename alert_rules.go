package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Rule string format: cpu.all > 50 : CPU

type AlertRule struct {
	Property string `json:"property"`
	Operator string `json:"operator"`
	Value    uint64 `json:"value"`
	name     string
}

func parseRuleFile(filePath string) ([]*AlertRule, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open alert rule file [%s], recieved error: %s", filePath, err)
	}
	defer file.Close()

	rules := []*AlertRule{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || line[0:1] == "#" {
			continue
		}

		rule, err := parseRuleString(line)
		if err != nil {
			return nil, fmt.Errorf("error parsing rule file, %s", err)
		}

		rules = append(rules, rule)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error parsing rule file, %s", err)
	}

	return rules, nil
}

func parseRuleString(ruleStr string) (*AlertRule, error) {
	// A placeholder is used here to track escaped colons, which are then restored
	// when the logic and name are parsed out.
	placeholder := "~~~|~~~"
	ruleStrEscapeMod := strings.ReplaceAll(ruleStr, `\:`, placeholder)
	ruleNameSplit := strings.Split(strings.TrimSpace(ruleStrEscapeMod), ":")
	ruleLogic := strings.ReplaceAll(strings.TrimSpace(ruleNameSplit[0]), placeholder, ":")

	ruleName := ruleLogic
	if len(ruleNameSplit) > 1 {
		ruleName = strings.ReplaceAll(strings.TrimSpace(ruleNameSplit[1]), placeholder, ":")
	}

	logicSplit := strings.Split(ruleLogic, " ")
	if len(logicSplit) != 3 {
		return nil, fmt.Errorf(`rule "%s" does not adhere to the format "<property> <operator> <value> : <name>"`, ruleStr)
	}

	validOperators := []string{">", "<", "=", "!=", "<=", ">="}
	if !contains(validOperators, logicSplit[1]) {
		return nil, fmt.Errorf(`operator in rule "%s" does not match allowed operators (%s)`, ruleStr, strings.Join(validOperators, " "))
	}

	property := logicSplit[0]
	operator := logicSplit[1]
	value, err := strconv.ParseInt(logicSplit[2], 10, 64)
	if err != nil {
		return nil, fmt.Errorf(`rule "%s" contained a non-numeric value of "%s"`, ruleStr, logicSplit[2])
	}

	rule := &AlertRule{
		name:     ruleName,
		Property: property,
		Operator: operator,
		Value:    uint64(value),
	}

	return rule, nil
}
