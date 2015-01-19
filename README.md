# fwdr

This library provides a simple router interface while also exposing the power of underlying regexes directly to the developer.

# Usage


```go

func main() {
	r := fwdr.NewRouter()

	r.Handler("/post/:id/:title", fwdr.NewReq("Id", "[0-9]+"), func(w http.ResponseWriter, r *http.Request) {
	
		vars := fwdr.Params(r)

		log.Printf("request id:%s title:%s", vars["id"], vars["title"])

		// logic here
	})
	
}

```

Overview

Parameters are used in the route to indicate how it should be parsed. 

In the following example route are indicated by `:title` and `:id` where test is the name used to look it up using `Params`.

```
	r.Handler("/post/:id/:title", fwdr.NewReq("Id", "[0-9]+"), func(w http.ResponseWriter, r *http.Request) {
```

Note: Parameters are case insensitive and are all converted to lower case internally, this is done to avoid confusion and typos. In the example `Id` and `id` result in the later being the param name.

Internally these parameters will be stored as: 

* `id` with a matcher of `([0-9]+)`
* `title` with a matcher of `(.+?)`

Parameters work as follows:

* they must be unique within a given route
* requirements are used to override the default pattern for a given parameter and aren't required

# Sponsor

This project was made possible by [Ninja Blocks](http://ninjablocks.com).

# License

This code is Copyright (c) 2014 Mark Wolfe and licenced under the MIT licence. All rights not explicitly granted in the MIT license are reserved. See the included LICENSE.md file for more details.