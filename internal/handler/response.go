package handler

import (
	"log"
	"net/http"
)

func newErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	log.Panic(message)
	w.WriteHeader(statusCode)
}
