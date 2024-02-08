package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bem-vindo ao meu webservice em Go!")
	})

	addr := ":8080"

	fmt.Printf("Servidor iniciado em http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}
