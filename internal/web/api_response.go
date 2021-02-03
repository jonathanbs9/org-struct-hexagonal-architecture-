package web

import (
	"encoding/json"
	"net/http"
)

// ResponseAPI struct
type ResponseAPI struct {
	Sucess bool        `json:"sucess"`
	Status int         `json:"status,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

// Success func
func Success(result interface{}, status int) *ResponseAPI {
	return &ResponseAPI{
		Sucess: true,
		Status: status,
		Result: result,
	}
}

// Send func
func (r *ResponseAPI) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	return json.NewEncoder(w).Encode(r)
}
