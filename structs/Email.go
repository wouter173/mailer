package structs

type Email struct {
	Target  string `json:"target"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
