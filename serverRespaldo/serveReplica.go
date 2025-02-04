package serverrespaldo

import (
	"fmt"
	"log"
	"net/http"
	"time"
)


func checkChanges() {
	for {
		resp, err := http.Get("http://localhost:8080/subscribe")
		if err != nil {
			fmt.Println("Error al consultar cambios:", err)
		} else if resp.StatusCode == http.StatusOK {
			fmt.Println("Se detectaron cambios. Actualizando...")
			updateReplica()
		}
		time.Sleep(5 * time.Second)
	}
}

func updateReplica() {
	resp, err := http.Get("http://localhost:8080/list")
	if err != nil {
		fmt.Println("Error al actualizar usuarios:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Usuarios actualizados recibidos")
}


func Run() {
	go checkChanges()

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Servidor de respaldo activo")
	})

	log.Println("Servidor de respaldo corriendo en :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
