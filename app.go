package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthorResp struct {
	Id      string    `json:"id"`
	Authors []*Author `json:"authors"`
}

type Author struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/health").Subrouter().Methods(http.MethodGet).Subrouter().HandleFunc("", health)
	r.PathPrefix("/authors/{productId}").Subrouter().Methods(http.MethodGet).Subrouter().HandleFunc("", bookAuthorsById)
	http.Handle("/", r)

	log.Println("Start listening http port 9080 ...")
	if err := http.ListenAndServe(":9080", nil); err != nil {
		panic(err)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	resp, err := json.Marshal(map[string]string{
		"status": "Authors is healthy",
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Write(resp)
}

func getAuthors(productId string) *AuthorResp {
	authorResp := &AuthorResp{
		Id: productId,
		Authors: []*Author{
			{
				Name: "KubeVela",
				Desc: "(26 April 1564 – 23 April 1616) was an English playwright, poet, and actor, widely regarded as the greatest writer in the English language and the world's greatest dramatist. He is often called England's national poet and the \"Bard of Avon\" (or simply \"the Bard\"). His extant works, including collaborations, consist of some 39 plays, 154 sonnets, three long narrative poems, and a few other verses, some of uncertain authorship. His plays have been translated into every major living language and are performed more often than those of any other playwright.",
			},
		},
	}

	return authorResp
}

func bookAuthorsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]

	authors := getAuthors(productId)
	resp, err := json.Marshal(authors)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Write(resp)
}
