package main

import (
	"encoding/json"
	"net/http"
)

// MemberIndex ...
func MemberIndex(w http.ResponseWriter, r *http.Request) {
	members := []Member{
		{ID: 1, Name: "name1"},
		{ID: 2, Name: "name2"},
	}

	if err := json.NewEncoder(w).Encode(members); err != nil {
		panic(err)
	}
}
