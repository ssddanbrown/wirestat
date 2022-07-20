package main

import "fmt"

func RunRulesAgainstMetrics(rules map[string]*AlertRule, metrics map[string]uint64) []string {
	var alerts []string

	for _, rule := range rules {
		fails, err := ruleFails(rule, metrics)
		if err != nil {
			alerts = append(alerts, err.Error())
			continue
		}

		if fails {
			alerts = append(alerts, rule.name)
		}
	}

	return alerts
}

func ruleFails(rule *AlertRule, metrics map[string]uint64) (bool, error) {

	val, exists := metrics[rule.Property]
	if !exists {
		return false, fmt.Errorf(`rule contains non-existing property "%s"`, rule.Property)
	}

	if rule.Operator == ">" {
		return val > rule.Value, nil
	}
	if rule.Operator == "<" {
		return val < rule.Value, nil
	}
	if rule.Operator == "=" {
		return val == rule.Value, nil
	}
	if rule.Operator == "!=" {
		return val != rule.Value, nil
	}
	if rule.Operator == "<=" {
		return val <= rule.Value, nil
	}
	if rule.Operator == ">=" {
		return val >= rule.Value, nil
	}

	return false, fmt.Errorf(`rule contains invalid operator "%s"`, rule.Operator)
}
