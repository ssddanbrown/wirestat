package main

import (
	"reflect"
	"strings"
)

func RunRulesAgainstSystem(rules map[string]*AlertRule, system *System) []string {
	var alerts []string

	for _, rule := range rules {
		if ruleFails(rule, system) {
			alerts = append(alerts, rule.name)
		}
	}

	return alerts
}

func ruleFails(rule *AlertRule, system *System) bool {

	val := getSystemValAtProp(rule.Property, system)

	if rule.Operator == ">" {
		return val > rule.Value
	}
	if rule.Operator == "<" {
		return val < rule.Value
	}
	if rule.Operator == "=" {
		return val == rule.Value
	}
	if rule.Operator == "!=" {
		return val != rule.Value
	}
	if rule.Operator == "<=" {
		return val <= rule.Value
	}
	if rule.Operator == ">=" {
		return val >= rule.Value
	}

	panic("Did not match an operator during rule check")
}

func getSystemValAtProp(prop string, system *System) uint64 {
	splitProp := strings.Split(prop, ".")
	var cObject interface{} = system
	var val uint64

	for i, jsonProp := range splitProp {
		if cObject == nil {
			break
		}

		atEnd := i == len(splitProp)-1
		objType := reflect.ValueOf(cObject).Kind()
		if objType == reflect.Map {
			mapValtype := reflect.TypeOf(cObject).Elem()

			// Flatten everything, All sub system checks so everything is a giant map
			// that we can easily access using this key.
			// Could be under a "metrics" key

			if atEnd {
				val = cObject[jsonProp].(uint64)
			} else {
				cObject = cObject.(map[string]any)[jsonProp]
			}

			continue
		}

		prop := jsonToPropCase(jsonProp)
		r := reflect.ValueOf(cObject)
		f := reflect.Indirect(r).FieldByName(prop)

		if atEnd {
			val = f.Uint()
		} else {
			cObject = f.Interface()
		}

	}

	return val
}

func jsonToPropCase(jsonProp string) string {
	propCase := ""
	parts := strings.Split(jsonProp, "_")
	for _, part := range parts {
		propCase += strings.ToUpper(part[0:1]) + part[1:]
	}
	return propCase
}
