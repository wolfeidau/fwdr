package fwdr

import (
	"net/http"
	"regexp"
	"strings"
)

const (
	paramChar      = ":"
	defaultMatcher = ".+?"
)

var (
	paramRegex = regexp.MustCompile(paramChar + "([a-zA-Z_]+)")
)

// NewRouter builds a new router.
func NewRouter() *Router {
	return &Router{}
}

// Router holds the state of all http routes and demuxes requests across them in the based order they are added.
type Router struct {

	// Configurable Handler to be used when there is no matching route.
	NotFoundHandler http.Handler

	// Routes to be matched, in order.
	routes []*Route
}

// ServeHTTP method is used as an entry point for the mux by the http.Server
func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var match routeMatch
	var handler http.Handler

	if rtr.Match(r, &match) {
		handler = match.handler
		//setParams(r, match.Params)
		//		setCurrentRoute(r, match.Route)
	}
	if handler == nil {
		handler = rtr.NotFoundHandler
		if handler == nil {
			handler = http.NotFoundHandler()
		}
	}

	handler.ServeHTTP(w, r)
}

func (rtr *Router) Match(r *http.Request, match *routeMatch) bool {

	for _, rt := range rtr.routes {
		matches := rt.regexp.FindAllStringSubmatch(r.URL.RequestURI(), -1)

		log.Debugf("matches %v", matches)

		if matches == nil {
			continue
		}

		match.handler = rt.handler
		match.params = make(map[string]string)

		for i, attr := range rt.regexp.SubexpNames() {
			if attr == "" {
				continue
			}
			match.params[attr] = matches[0][i]
		}

		// log.Debugf("match.params %v", match.params)

		return true
	}

	return false

}

// HandleFunc is used to attach a route to the router.
func (rtr *Router) HandleFunc(pattern string, requirements *Reqs, handler func(http.ResponseWriter, *http.Request)) {
	rtr.routes = append(rtr.routes, newRoute(requirements, pattern, handler))
}

// Reqs is the requirements of a route
type Reqs struct {
	reqs map[string]string
}

// NewReq this function creates the first requirement and puts it into a requirements struct
// which uses the builder pattern.
func NewReq(attr, matcher string) *Reqs {
	rs := &Reqs{reqs: make(map[string]string)}

	rs.addReq(attr, matcher)

	return rs
}

// NewReq this method enables you to make a new requirement and add it to the requirements struct, it supports chaining.
func (rs *Reqs) NewReq(attr, matcher string) *Reqs {

	rs.addReq(attr, matcher)

	return rs
}

func (rs *Reqs) addReq(attr, matcher string) {

	attr = strings.ToLower(attr)

	rs.reqs[attr] = matcher

	// assign the default matcher
	if matcher == "" {
		matcher = defaultMatcher
	}

}

func (rs *Reqs) mustValidate() {

	// validate the requirements
	for attr, matcher := range rs.reqs {

		log.Debugf("validating %s (%s)", attr, matcher)

		_, err := regexp.Compile(matcher)

		if err != nil {
			panic(err)
		}
	}
}

// routeMatch route match
type routeMatch struct {
	handler http.Handler
	params  map[string]string
}
