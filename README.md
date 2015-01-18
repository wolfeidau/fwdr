# fwdr

This library provides a simple router interface while also exposing the power of underlying regexes directly to the developer.

# Usage


```go

func main() {
	r := fwdr.NewRouter()

	r.Handler("/post/:id/:title", fwdr.NewReq("Id", "[0-9]+"), func(w http.ResponseWriter, r *http.Request) {
	
		vars := fwdr.Vars(r)

		log.Printf("request id:%s title:%s", vars["id"], vars["title"])

		// logic here
	})
	
}

```

Overview

Parameters are used in the route to indicate how it should be parsed. 

In the following example route are indicated by `:title` and `:Id` where test is the name used to look it up using `Params`.

```
	r.Handler("/post/:Id/:title", fwdr.NewReq("Id", "[0-9]+"), func(w http.ResponseWriter, r *http.Request) {
```

Internally these parameters will be stored as: 

* `id` with a matcher of `([0-9]+)`
* `title` with a matcher of `(.+?)`

Parameters work as follows:
* they are case insensitive, the are all converted to lower case internally, this is done to avoid confusion and typos
* they must be unique within a given route
* requirements are used to override the default pattern for a given parameter and aren't required


