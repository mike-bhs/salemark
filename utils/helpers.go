package utils

import (
	"net/http"
	"strings"
)

func RequestParams(req *http.Request) map[string]string {
	m := make(map[string]string)

	for key, values := range req.Form {
		v := strings.Join(values[:], ", ")
		m[key] = v
	}

	return m
}
