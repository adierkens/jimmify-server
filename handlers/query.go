package handlers

import (
	"encoding/json"
	"jimmify-server/db"
	"net/http"
)

//Query: submit a query
func Query(w http.ResponseWriter, r *http.Request) {
	var q db.Query
	response := make(map[string]interface{})

	//read json
	err := json.NewDecoder(r.Body).Decode(&q)
	if err != nil {
		ReturnStatusBadRequest(w, "Failed to decode query json")
		return
	}

	//add query
	key, err := db.AddQuery(q)
	if err != nil {
		ReturnInternalServerError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response["key"] = key
	response["status"] = "true"
	json.NewEncoder(w).Encode(response)
}