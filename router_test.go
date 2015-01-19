package fwdr

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/juju/loggo"
)

var tlog = loggo.GetLogger("test")

func TestBuildRouter(t *testing.T) {
	SetLogger(tlog)

	tlog.SetLogLevel(loggo.DEBUG)

	r := NewRouter()

	r.HandleFunc("/post/:id/:post_title", NewReq("Id", "[0-9]+"), func(w http.ResponseWriter, r *http.Request) {
		// logic here
	})

	http.Handle("/", r)
}

func TestRoute(t *testing.T) {

	req, _ := http.NewRequest("GET", "/post/123/testing", nil)
	w := httptest.NewRecorder()

	r := NewRouter()

	r.HandleFunc("/post/:id/:post_title", NewReq("Id", "[0-9]+"), func(w http.ResponseWriter, r *http.Request) {
		// logic here
		log.Debugf("r %s", r.URL.String())

		w.WriteHeader(http.StatusOK)
	})

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("route didn't return %v", http.StatusOK)
	}

}
