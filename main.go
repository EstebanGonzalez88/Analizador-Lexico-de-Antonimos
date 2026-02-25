package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type Persona struct {
	Nombre  string `json:"nombre"`
	Paterno string `json:"paterno"`
	Materno string `json:"materno"`
	Fecha   string `json:"fecha"`
	Sexo    string `json:"sexo"`
	Estado  string `json:"estado"`
}

func primeraVocal(s string) string {
	for i := 1; i < len(s); i++ {
		if strings.ContainsRune("AEIOU", rune(s[i])) {
			return string(s[i])
		}
	}
	return "X"
}

func primeraConsonante(s string) string {
	for i := 1; i < len(s); i++ {
		if !strings.ContainsRune("AEIOU", rune(s[i])) {
			return string(s[i])
		}
	}
	return "X"
}
func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.HandleFunc("/curp", generarCURPHandler)

	fmt.Println("Servidor en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func generarCURPHandler(w http.ResponseWriter, r *http.Request) {
	var p Persona

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Error en datos", http.StatusBadRequest)
		return
	}

	curp := generarCURP(p)

	response := map[string]string{"curp": curp}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func generarCURP(p Persona) string {
	nombre := strings.ToUpper(strings.TrimSpace(p.Nombre))
	paterno := strings.ToUpper(strings.TrimSpace(p.Paterno))
	materno := strings.ToUpper(strings.TrimSpace(p.Materno))

	fecha, _ := time.Parse("2006-01-02", p.Fecha)

	// Iniciales básicas
	curp := ""
	curp += string(paterno[0])
	curp += primeraVocal(paterno)
	curp += string(materno[0])
	curp += string(nombre[0])

	// Fecha YYMMDD
	curp += fecha.Format("060102")

	// Sexo
	curp += strings.ToUpper(p.Sexo)

	// Estado (solo SLP o CAM)
	curp += strings.ToUpper(p.Estado)

	// Consonantes internas
	curp += primeraConsonante(paterno)
	curp += primeraConsonante(materno)
	curp += primeraConsonante(nombre)

	// Homoclave simulada
	curp += "00"

	return curp
}
