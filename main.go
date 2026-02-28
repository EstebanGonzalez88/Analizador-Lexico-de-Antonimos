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

func main() {

	// Servir carpeta public
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.HandleFunc("/curp", generarCURP)

	fmt.Println("Servidor en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generarCURP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var p Persona
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	// Validar fecha
	fecha, err := time.Parse("2006-01-02", p.Fecha)
	if err != nil {
		http.Error(w, "Fecha inválida", http.StatusBadRequest)
		return
	}

	hoy := time.Now()
	if fecha.After(hoy) {
		http.Error(w, "Fecha futura no permitida", http.StatusBadRequest)
		return
	}

	// Validar edad máxima 150
	edad := hoy.Year() - fecha.Year()
	if hoy.YearDay() < fecha.YearDay() {
		edad--
	}
	if edad > 150 {
		http.Error(w, "Edad mayor a 150 años no permitida", http.StatusBadRequest)
		return
	}

	if p.Estado != "SLP" && p.Estado != "CAM" {
		http.Error(w, "Entidad no permitida", http.StatusBadRequest)
		return
	}

	if p.Sexo != "H" && p.Sexo != "M" {
		http.Error(w, "Sexo inválido", http.StatusBadRequest)
		return
	}

	curp := construirCURP(p, fecha)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"curp": curp,
	})

	// validar que el nombre teng Minimo 3 caracteres
	if len(p.Nombre) < 3 {
		http.Error(w, "Nombre debe tener al menos 3 caracteres", http.StatusBadRequest)
		return
	}
	// Validar que el apellido paterno tenga minimo 3 caracteres
	if len(p.Paterno) < 3 {
		http.Error(w, "Apellido paterno debe tener al menos 3 caracteres", http.StatusBadRequest)
		return
	}
	// Validar que el apellido materno tenga minimo 3 caracteres
	if len(p.Materno) < 3 {
		http.Error(w, "Apellido materno debe tener al menos 3 caracteres", http.StatusBadRequest)
		return
	}
	// Validar que el nombre no contenga caracteres especiales
	if strings.ContainsAny(p.Nombre, "!@#$%^&*()_+={}[]|\\:;\"'<>,.?/") {
		http.Error(w, "Nombre no debe contener caracteres especiales", http.StatusBadRequest)
		return
	}
	// Validar que el apellido paterno no contenga caracteres especiales
	if strings.ContainsAny(p.Paterno, "!@#$%^&*()_+={}[]|\\:;\"'<>,.?/") {
		http.Error(w, "Apellido paterno no debe contener caracteres especiales", http.StatusBadRequest)
		return
	}
	// Validar que el apellido materno no contenga caracteres especiales
	if strings.ContainsAny(p.Materno, "!@#$%^&*()_+={}[]|\\:;\"'<>,.?/") {
		http.Error(w, "Apellido materno no debe contener caracteres especiales", http.StatusBadRequest)
		return
	}
	// Validar que el nombre no contenga números
	if contieneNumeros(p.Nombre) {
		http.Error(w, "Nombre no debe contener números", http.StatusBadRequest)
		return
	}
	// Validar que el apellido paterno no contenga números
	if contieneNumeros(p.Paterno) {
		http.Error(w, "Apellido paterno no debe contener números", http.StatusBadRequest)
		return
	}
	// Validar que el apellido materno no contenga números
	if contieneNumeros(p.Materno) {
		http.Error(w, "Apellido materno no debe contener números", http.StatusBadRequest)
		return
	}

}

func construirCURP(p Persona, fecha time.Time) string {

	nombre := strings.ToUpper(strings.TrimSpace(p.Nombre))
	paterno := strings.ToUpper(strings.TrimSpace(p.Paterno))
	materno := strings.ToUpper(strings.TrimSpace(p.Materno))

	curp := ""

	curp += string(paterno[0])
	curp += primeraVocal(paterno)
	curp += string(materno[0])
	curp += string(nombre[0])

	curp += fecha.Format("060102")
	curp += p.Sexo
	curp += p.Estado

	curp += primeraConsonante(paterno)
	curp += primeraConsonante(materno)
	curp += primeraConsonante(nombre)

	curp += "00"

	return curp
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

// Mostrar mensaje de error si el nombre, apellido paterno o apellido materno contienen números
func contieneNumeros(s string) bool {
	for _, char := range s {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}
