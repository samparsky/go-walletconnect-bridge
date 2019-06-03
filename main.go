package main
import (
 "log"
 "fmt"
 "net/http"
 "github.com/gorilla/mux"
 "encoding/json"
)

var version = "1.0.0-beta"
var addr = ":8001"
var description = "walletconnect golang implementation"
var name = "go-walletconnect"

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloHandler)
	router.HandleFunc("/info", infoHandler)
	router.HandleFunc("/health", healthHandler)
	router.HandleFunc("/subscribe", subscribeHandler).Methods("POST")
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal("failed to start server")
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Hello World, this is WalletConnect %s", version)
	jsonResponse(w, []byte(response))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	w.Write(nil)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Name string
		Version string
		Description string
	}{
		name,
		version,
		description,
	}

	response, _ := json.Marshal(data)
	jsonResponse(w, response)
}

func subscribeHandler(w http.ResponseWriter, r *http.Request) {
	type Body struct {
		Topic string
		Webhook string
	}
	body := Body{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Printf("Failed to parse json body %v", err)
		errorResponse(w, "Failed to parse json body")
		return
	}
	if body.Topic == "" {
		errorResponse(w, "Error: missing or invalid topic field")
		return
	}
	if body.Webhook == "" {
		errorResponse(w, "Error: missing or invalid webhook field")
		return
	}

	data := struct {
		success bool
	}{
		true,
	}
	response, _ := json.Marshal(data)
	jsonResponse(w, response)
}

func jsonResponse(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func errorResponse(w http.ResponseWriter, err string){
	data := struct {
		Error bool
		Message string
	}{
		true,
		err,
	}
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(response)
}