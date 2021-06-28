package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Handler exports http.HandlerFuncs that provide the main functionality behind
// your REST APIs.
type Handler struct {
	*mux.Router
}

// New is the constructor for the Handler type.
func (h *Handler) New() error {
	h.Router = mux.NewRouter()
	h.configureRouter()
	return nil
}

func (h *Handler) configureRouter() {
	r := h.Router

	r.HandleFunc("/search/{value}", h.search).Methods(http.MethodGet, http.MethodOptions)
}

// IssueMetadata ...
type IssueMetadata struct {
	Title  *string        `json:"title"`
	Labels []github.Label `json:"labels"`
}

// SearchResponse ...
type SearchResponse struct {
	Count      int            `json:"total_count"`
	Incomplete bool           `json:"incomplete_results"`
	Issues     []github.Issue `json:"items"`
}

func setupHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// search returns search results for issues in facebook/react repo
func (h *Handler) search(w http.ResponseWriter, r *http.Request) {
	setupHeader(w)
	params := mux.Vars(r)
	value, ok := params["value"]
	if !ok {
		http.Error(w, "missing route param value", http.StatusBadRequest)
		return
	}

	q := url.Values{}
	q.Add("q", "repo:facebook/react "+value)

	query := "https://api.github.com/search/issues?" + q.Encode()
	log.Info(query)
	resp, err := http.Get(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting issues, ERROR: %v", err), http.StatusInternalServerError)
		return
	}

	log.Info(resp)
	issues := SearchResponse{}

	err = json.NewDecoder(resp.Body).Decode(&issues)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding issues, ERROR: %v", err), http.StatusInternalServerError)
		return
	}

	metadata := []IssueMetadata{}
	for _, issue := range issues.Issues {
		m := IssueMetadata{}
		m.Title = issue.Title
		m.Labels = issue.Labels
		metadata = append(metadata, m)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(metadata)
}
