package fwdr

import "net/http"

// Request which wraps the standard http.Request and adds params to it
// to store the route params.
type Request struct {
	*http.Request
	params map[string]string
}
