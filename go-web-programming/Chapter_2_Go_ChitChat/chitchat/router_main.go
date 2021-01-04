package main

import "net/http"

// Get /err?msg=
// show the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	}
}
