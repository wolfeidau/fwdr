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
	paramRegex = regexp.MustCompile(paramChar + "([a-z_]+)")
)

// NewRouter builds a new router.
func NewRouter() *Router {
	return &Router{}
}

// Router holds the state of all http routes and demuxes requests across them in the based order they are added.
type Router struct {
	// Routes to be matched, in order.
	routes []*Route
}

// ServeHTTP method is used as an entry point for the mux by the http.Server
func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

// HandleFunc is used to attach a route to the router.
func (rtr *Router) HandleFunc(pattern string, requirements *Reqs, handler func(http.ResponseWriter, *http.Request)) {

	rt := &Route{rs: requirements, pattern: pattern}

	rt.assignParams()

	rtr.routes = append(rtr.routes, rt)

}

// Route holds the state for a route.
type Route struct {
	// requirements for this route
	rs *Reqs

	pattern string

	regexp *regexp.Regexp
}

func (rt *Route) assignParams() {

	// build a list of requirements based on the route
	mtchs := paramRegex.FindAllStringSubmatch(rt.pattern, -1)

	log.Debugf("matches %v", mtchs)

	// iterate over the matching pairs to fill in those which didn't have specific
	// requirements specified
	for _, mt := range mtchs {
		attr := strings.ToLower(mt[1])

		if _, ok := rt.rs.reqs[attr]; !ok {
			rt.rs.reqs[attr] = defaultMatcher
		}
	}

	// TODO do the reverse of the above and locate requirements which don't match a
	// specific attribute in the route.

	// validate
	rt.rs.mustValidate()

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
