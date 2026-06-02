package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
)

func calculadoraHandler(w http.ResponseWriter, r *http.Request) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")
	op := r.URL.Query().Get("op")

	a, err1 := strconv.ParseFloat(aStr, 64)
	b, err2 := strconv.ParseFloat(bStr, 64)

	if err1 != nil || err2 != nil {
		http.Error(w, "Uso: ?a=5&b=3&op=suma (Opciones: suma, resta, multiplica, divide)", http.StatusBadRequest)
		return
	}

	var res float64
	switch op {
	case "suma": res = a + b
	case "resta": res = a - b
	case "multiplica": res = a * b
	case "divide":
		if b == 0 { http.Error(w, "No se puede dividir por 0", http.StatusBadRequest); return }
		res = a / b
	default:
		http.Error(w, "Operación no válida", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Resultado: %.2f\n", res)
}

func passwordHandler(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 12)
	rand.Read(b)
	fmt.Fprintf(w, "Password seguro: %s\n", base64.URLEncoding.EncodeToString(b))
}

func main() {
	http.HandleFunc("/calcular", calculadoraHandler)
	http.HandleFunc("/password", passwordHandler)
	fmt.Println("Servidor escuchando en el puerto 8080...")
	http.ListenAndServe(":8080", nil)
}