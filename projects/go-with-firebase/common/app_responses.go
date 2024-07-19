package common

type SuccessResponse struct {
	Data  interface{} `json:"data"` //interface{}: means any data type
	Error error       `json:"error"`
}

type IntentResponse struct {
	ClientSecret string `json:"client-secret"`
}
