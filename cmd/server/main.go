package main

import (
	gordleHttp "github.com/TheDonDope/gordle/pkg/http"
)

func main() {

	r := gordleHttp.NewServer()
	err := r.Run(":4242")
	if err != nil {
		panic(err)
	}

}
