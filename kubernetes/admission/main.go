package main

import "fmt"

func serveMutatePods(w http.ResponseWriter, r *http.Request) {
	serve(w, r, mutatePods)
}

func main() {
	var config Config
	config.addFlags()
	flag.Parse()

	http.HandleFunc("/mutating-pods", serveMutatePods)
}
