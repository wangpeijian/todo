package core

import (
	"net/http"
)

type RequireHandle struct {
	Request *http.Request
	Body []byte
	Form map[string][]string
	Query map[string][]string
	JsonForm map[string] interface{}
}