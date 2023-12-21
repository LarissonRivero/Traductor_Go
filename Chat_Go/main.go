package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/translate/v2"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func handleTranslation(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		enableCors(&w)
		return
	}

	text := r.FormValue("text")
	targetLanguage := r.FormValue("targetLanguage")

	if targetLanguage != "es" && targetLanguage != "en" {
		http.Error(w, "Idioma de destino no válido", http.StatusBadRequest)
		return
	}

	apiKey := "AIzaSyDBSTZPv3_ENbI7Ty-NkcQkBVIwEgGfkls"
	ctx := context.Background()
	client, err := translate.NewService(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		log.Printf("Error al crear el cliente: %v", err)
		http.Error(w, "Error al crear el cliente", http.StatusInternalServerError)
		return
	}

	resp, err := client.Translations.List([]string{text}, targetLanguage).Do()
	if err != nil {
		log.Printf("Error al realizar la traducción: %v", err)
		http.Error(w, "Error al realizar la traducción", http.StatusInternalServerError)
		return
	}

	translatedText := resp.Translations[0].TranslatedText

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"original": text, "translation": translatedText})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/translate", handleTranslation).Methods("POST", "OPTIONS")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
