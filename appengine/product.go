// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"
)

type Product struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		Product{Id: "000001", Name: "Computer"},
		Product{Id: "000002", Name: "Smartphone"},
		Product{Id: "000003", Name: "Keyboard"},
	}

	json.NewEncoder(w).Encode(map[string][]Product{"products": products})
}
