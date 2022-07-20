package main

type ResponseBuilder struct {
	rules map[string]*AlertRule
}

type Response struct {
	Alerts  []string              `json:"alerts"`
	Rules   map[string]*AlertRule `json:"rules"`
	Metrics map[string]uint64     `json:"metrics"`
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
	metricMap := system.mergeMaps()
	alerts := RunRulesAgainstMetrics(r.rules, metricMap)

	return &Response{
		Metrics: metricMap,
		Rules:   r.rules,
		Alerts:  alerts,
	}
}
