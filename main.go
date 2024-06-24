package main

import (
	"handlers"
	"log"
	"net/http"
	"time"
)

type Info struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	mux.Handle("/download/", http.StripPrefix("/download/", http.FileServer(http.Dir("./web/download"))))

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/about", handlers.About)

	server := &http.Server{
		Addr:              "localhost:8080", // adresse du server (le port choisi est à titre d'exemple)
		Handler:           mux,              // listes des handlers
		ReadHeaderTimeout: 10 * time.Second, // temps autorisé pour lire les headers
		WriteTimeout:      10 * time.Second, // temps maximum d'écriture de la réponse
		IdleTimeout:       10 * time.Second, // temps maximum entre deux rêquetes
		ReadTimeout:       20 * time.Second, // durée maximale autorisée pour lire la requête complète, y compris le corps de la requête.
		MaxHeaderBytes:    5 << 20,          // 1 MB // maximum de bytes que le serveur va lire
	}

	handlers.GetFromJSON("https://groupietrackers.herokuapp.com/api/artists", &handlers.VarArtists)
	handlers.GetFromJSON("https://groupietrackers.herokuapp.com/api/relation", &handlers.VarDateLocation)

	log.Printf("Server starting on http://%s...\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
