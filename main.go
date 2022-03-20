package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/taigrr/sidecar-server/exe"
	"github.com/taigrr/sidecar-server/types"
)

func main() {
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
func PostTabs(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	tabList := types.TabSet{}
	err := json.NewDecoder(r.Body).Decode(&tabList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	a := types.Actions{}
	for _, e := range tabList.Tabs {
		fmt.Printf("Checking URL: %s\n", e.URL)
		shouldClose, err := exe.Spawn(e.URL)
		if err != nil {
			fmt.Printf("Error spawning action: %v\n", err)
		}
		if shouldClose {
			a.Close = append(a.Close, e.ID)
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(a)

	//update response writer
}
