package types

import "time"

type Issue struct {
	Code           string   `json:"code"`
	Message        string   `json:"message"`
	Severity       string   `json:"severity"`
	AttributeNames []string `json:"attributeNames"`
	Categories     []string `json:"categories"`
	Enforcements   []IssueEnforcements
}

type IssueEnforcements struct {
	Actions   []IssueEnforcementAction `json:"actions"`
	Exemption IssueExemption
}

type IssueEnforcementAction struct {
	Action string `json:"action"`
}

type IssueExemption struct {
	Status     string    `json:"status"`
	ExpiryDate time.Time `json:"expiryDate"`
}
