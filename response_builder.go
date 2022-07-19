package main

type ResponseBuilder struct {
	rules map[string]*AlertRule
}

type Response struct {
	Rules  map[string]*AlertRule `json:"rules"`
	Alerts []string              `json:"alerts"`
	*System
}

func NewResponseBuilder(rules []*AlertRule) *ResponseBuilder {

	ruleMap := make(map[string]*AlertRule)
	for _, rule := range rules {
		ruleMap[rule.name] = rule
	}

	return &ResponseBuilder{rules: ruleMap}
}

func (r *ResponseBuilder) GetResponseData() *Response {
	system := GetLatestSystem()

	return &Response{
		System: system,
		Rules:  r.rules,
		Alerts: []string{},
	}
}
