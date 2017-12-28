package main 

import (
		"net/http"
)

type Route struct {
	Name		string
	Method		string
	Pattern		string
	HandlerFunc	http.HandlerFunc
}

type Routes []Route

var routes = Routes {
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"CodesIndex",
		"GET",
		"/codes",
		CodesIndex,
	},
	Route{
		"CodeShow",
		"GET",
		"/codes/{codeId}",
		CodeShow,
	},
	Route{
		"CodeCreate",
		"POST",
		"/codes",
		CodeCreate,
	},
	Route{
		"ModuleShow",
		"GET",
		"/module/{module}",
		ModuleShow,
	},
	Route{
		"FieldShow",
		"GET",
		"/field/{field}",
		FieldShow,
	},
}