package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Topic interface {
	Retrieve(id int) (err error)
	Create() (err error)
	Update() (err error)
	Delete() (err error)
}

func HandleRequest(t Topic) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var err error
		switch request.Method {
		case "GET":
			err = showPost(writer, request, t)
		case "POST":
			err = createPost(writer, request, t)
		case "PUT":
			err = updatePost(writer, request, t)
		case "DELETE":
			err = deletePost(writer, request, t)
		}
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func showPost(w http.ResponseWriter, r *http.Request, t Topic) (err error) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		return
	}
	err = t.Retrieve(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func createPost(w http.ResponseWriter, r *http.Request, t Topic) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, t)
	err = t.Create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func updatePost(w http.ResponseWriter, r *http.Request, t Topic) (err error) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		return
	}
	err = t.Retrieve(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, t)
	err = t.Update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func deletePost(w http.ResponseWriter, r *http.Request, t Topic) (err error) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		return
	}
	err = t.Retrieve(id)
	if err != nil {
		return
	}
	err = t.Delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
