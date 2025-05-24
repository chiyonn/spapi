package model

type ErrorList struct {
	Code    string `json:"code"`
	Message *string `json:"message"`
	Details *string `json:"details"`
}

