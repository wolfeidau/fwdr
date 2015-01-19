package fwdr

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

// Route holds the state for a route.
type Route struct {
	// requirements for this route
	rs *Reqs

	pattern string

	regexp *regexp.Regexp

	handler http.Handler
}

func newRoute(requirements *Reqs, pattern string, handler func(http.ResponseWriter, *http.Request)) *Route {

	rt := &Route{rs: requirements, pattern: pattern}

	rt.assignParams()
	rt.buildRegex()

	// assign the handler
	rt.handler = http.HandlerFunc(handler)

	return rt
}

func (rt *Route) assignParams() {

	// build a list of requirements based on the route
	mtchs := paramRegex.FindAllStringSubmatch(rt.pattern, -1)

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

func (rt *Route) buildRegex() {
	rp := strings.Replace(fmt.Sprintf("^%s$", rt.pattern), ".", "\\.", -1)

	for attr, matcher := range rt.rs.reqs {
		rp = strings.Replace(rp, paramChar+attr, fmt.Sprintf("(?P<%s>%s)", attr, matcher), 1)
	}

	//log.Debugf("final regex = %s", rp)

	rt.regexp = regexp.MustCompile(rp)
}
