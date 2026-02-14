package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Servir archivos desde el directorio actual
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
