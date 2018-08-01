package handler

import (
	"encoding/json"
	"net/http"

	"../model"
)

// MemberIndex ...
func MemberIndex(w http.ResponseWriter, r *http.Request) {
	members := []model.Member{
		{ID: 1, Name: "name1"},
		{ID: 2, Name: "name2"},
	}

	if err := json.NewEncoder(w).Encode(members); err != nil {
		panic(err)
	}
}
