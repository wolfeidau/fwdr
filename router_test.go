package fwdr

import (
	"net/http"
	"testing"

	"github.com/juju/loggo"
)

var tlog = loggo.GetLogger("test")

func TestBuildRouter(t *testing.T) {
	SetLogger(tlog)

	tlog.SetLogLevel(loggo.DEBUG)

	r := NewRouter()

	r.HandleFunc("/post/:id/:title", NewReq("Id", "[0-9]+"), func(w http.ResponseWriter, r *http.Request) {
		// logic here
	})

	http.Handle("/", r)
}
