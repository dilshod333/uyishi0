package api

import (
	"conn/service"
	"log"
	"net/http"

	_ "conn/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func ConnectApi() {
	router := http.NewServeMux()
	
	// Swagger UI
	router.Handle("/swagger/", httpSwagger.WrapHandler)
	
	// Create a new server instance
	n, err := service.NewServer()
	if err != nil {
		log.Println("Error creating server:", err)
		return
	}
	
	// Define routes
	router.HandleFunc("/books", n.CreateBook)
	router.HandleFunc("/books/", n.GetBookID)
	router.HandleFunc("/books/", n.UpdateBookk)
	router.HandleFunc("/books/", n.DeleteBookById)

	// Apply CORS middleware
	handler := enableCORS(router)

	log.Println("Running on :7777")
	log.Fatal(http.ListenAndServe(":7777", handler))
}
