package response

import (
	"fmt"
	"net/http"
)

// ErrorWithText prints out an error message to the user
func ErrorWithText(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

// SuccessWithJSON prints out a success message to the user, formated in JSON
func SuccessWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}