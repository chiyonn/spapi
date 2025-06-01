package types

type Status struct {
	StatusCode *int `json:"statusCode"`
	ReasonPhrase *string `json:"reasonPhrase"`
}
