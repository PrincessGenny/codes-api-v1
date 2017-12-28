package main 

import (
		"fmt"
		"net/http"
		"encoding/json"
		"io"
		"io/ioutil"
		"strconv"

		"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func CodesIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(codes); err != nil {
		panic(err)
	}
}

func CodeShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var codeId int
	var err error
	if codeId, err = strconv.Atoi(vars["codeId"]); err != nil {
		panic(err)
	}
	code := RepoFindCode(codeId)
	if code.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(code); err != nil {
			panic(err)
		}
		return
	}

	//404 if not found
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not found"}); err != nil {
		panic(err)
	}
}

func ModuleShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var module string
	var err error
	if module = vars["module"]; err != nil {
		panic(err)
	}
	codes := RepoFindCodeByModule(module)
	if len(codes) > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(codes); err != nil {
			panic(err)
		}
		return
	}

	//404 if not found
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not found"}); err != nil {
		panic(err)
	}
}

func CodeCreate(w http.ResponseWriter, r *http.Request) {
	var code Code
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body,&code); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}	
	}

	c := RepoCreateCode(code)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(c); err != nil {
		panic(err)
	}
}