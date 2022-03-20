package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func test() {
	r := mux.NewRouter()
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	r.HandleFunc("/route", PostTabs).Methods("POST", "OPTIONS")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "hello")
	})

	err := http.ListenAndServe(":15150", handlers.CORS(header, methods, origins)(r))

	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	//test()
	router := mux.NewRouter()

	//specify endpoints, handler functions and HTTP method
	router.HandleFunc("/post-tabs", PostTabs).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/health-check", HealthCheck).Methods(http.MethodGet, http.MethodOptions)
	http.Handle("/", router)

	//start and listen to requests
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"chrome-extension://*", "*", "http://localhost"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// start server listen
	// with error handling
	//	log.Fatal(http.ListenAndServe(":15150", router))
	log.Fatal(http.ListenAndServe(":15150", handlers.CORS(headersOk, originsOk, methodsOk)(router)))
}
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	//update response writer
	fmt.Fprintf(w, "{\"ok\": true}")
}

type TabSet struct {
	Tabs []Tab `json:"tabs"`
}
type MutedInfo struct {
	IsMuted bool `json:"muted"`
}

type Tab struct {
	IsActive          bool      `json:"active"`
	IsAudible         bool      `json:"audible"`
	IsAutoDiscardable bool      `json:"autoDiscardable"`
	IsDiscarded       bool      `json:"discarded"`
	FaviconURL        string    `json:"favIconUrl"`
	GroupID           int       `json:"groupId"`
	Height            int       `json:"height"`
	IsHighlighted     bool      `json:"highlighted"`
	ID                int       `json:"id"`
	IsIncognito       bool      `json:"incognito"`
	Index             int       `json:"index"`
	MutedInfo         MutedInfo `json:"mutedInfo"`
	IsPinned          bool      `json:"pinned"`
	IsSelected        bool      `json:"selected"`
	LoadedStatus      string    `json:"status"`
	Title             string    `json:"title"`
	URL               string    `json:"url"`
	Width             int       `json:"width"`
	WindowID          int       `json:"windowId"`
}

type Actions struct {
	Close []int `json:"close"`
}

func PostTabs(w http.ResponseWriter, r *http.Request) {
	tabList := TabSet{}

	err := json.NewDecoder(r.Body).Decode(&tabList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	a := Actions{}
	for i, e := range tabList.Tabs {
		fmt.Println(e.URL)
		if i == 0 {
			a.Close = []int{e.ID}
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(a)

	//update response writer
}
