package server

import (
	"Practica/clienteServidor/server/app"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var (
	arbol       = &app.ArbolBinario{}
	subscribers = make([]chan bool, 0)
	mu          sync.Mutex
)

func createPersonHandler(w http.ResponseWriter, r *http.Request) {
	var p app.Persona
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Error en el JSON", http.StatusBadRequest)
		return
	}

	mu.Lock()
	arbol.Insertar(p)
	mu.Unlock()

	notifySubscribers()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func deletePersonHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	mu.Lock()
	arbol.Eliminar(id)
	mu.Unlock()

	notifySubscribers()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Usuario con ID %d eliminado", id)
}

func subscribeChangesHandler(w http.ResponseWriter, r *http.Request) {
	notify := make(chan bool)
	mu.Lock()
	subscribers = append(subscribers, notify)
	mu.Unlock()

	<-notify
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Cambios detectados")
}

func notifySubscribers() {
	mu.Lock()
	defer mu.Unlock()
	for _, ch := range subscribers {
		ch <- true
	}
	subscribers = nil
}

func listPersonsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	personas := arbol.ObtenerTodos()
	json.NewEncoder(w).Encode(personas)
}

func Run() {
    http.HandleFunc("/create", createPersonHandler)
    http.HandleFunc("/delete", deletePersonHandler)
    http.HandleFunc("/subscribe", subscribeChangesHandler)
	http.HandleFunc("/list", listPersonsHandler) 

    log.Println("Servidor principal corriendo en :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
