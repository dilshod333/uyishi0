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
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080") // Restrict to Swagger UI
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

// @title Example API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /
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
	router.HandleFunc("POST /books/", n.CreateBook)
	router.HandleFunc("GET /books/{id}", n.GetBookID)
	router.HandleFunc("PUT /books/{id}", n.UpdateBookk)
	router.HandleFunc("DELETE /books/{id}", n.DeleteBookById)

	// Apply CORS middleware
	handler := enableCORS(router)

	log.Println("Running on :7777")
	log.Fatal(http.ListenAndServe(":7777", handler))
}
