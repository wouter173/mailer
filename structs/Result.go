package structs

import "encoding/json"

type Result struct {
	Error string `json:"error,omitempty"`
	Ok    bool   `json:"ok"`
}

//Create a new json marshalled result from the result struct.
func NewResult(e string) []byte {
	var r Result

	if e != "" {
		r = Result{Ok: false, Error: e}
	} else {
		r = Result{Ok: true}
	}

	res, _ := json.Marshal(r)
	return res
}
