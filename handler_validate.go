package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	type returnVals struct {
		Valid bool `json:"valid"`
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		errResp := errorResponse{
			Error: "Something went wrong",
		}

		data, err := json.Marshal(errResp)
		if err != nil {
			log.Printf("marshalling error JSON: %s", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write(data)
		return
	}

	if len(params.Body) > 140 {
		errResp := errorResponse{
			Error: "Chirp is too long",
		}

		data, err := json.Marshal(errResp)
		if err != nil {
			log.Printf("marshalling error JSON: %s", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(data)
		return
	}

	respBody := returnVals{Valid: true}
	data, err := json.Marshal(respBody)
	if err != nil {
		errResp := errorResponse{
			Error: "Something went wrong",
		}

		data, err := json.Marshal(errResp)
		if err != nil {
			log.Printf("marshalling error JSON: %s", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write(data)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Printf("vaild:%v", respBody)
	w.WriteHeader(200)
	w.Write(data)

}
