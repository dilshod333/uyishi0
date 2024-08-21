package api

import (
	"conn/service"
	"expvar"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "conn/docs" 

	"github.com/felixge/httpsnoop"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Example API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /
func ConnectApi() {
	router := http.NewServeMux()
	router.Handle("/swagger/", httpSwagger.WrapHandler)
	n, err := service.NewServer()
	if err != nil {
		log.Println("Xatolik", err)
		return
	}
	create := http.HandlerFunc(n.CreateBook)
	gett := http.HandlerFunc(n.GetBookID)
	updatee := http.HandlerFunc(n.UpdateBookk)
	deletee := http.HandlerFunc(n.DeleteBookById)
	router.HandleFunc("POST /books", (create))
	router.HandleFunc("GET /books/{id}", gett)
	router.HandleFunc("PUT /books/{id}", updatee)
	router.HandleFunc("DELETE /books/{id}", deletee)
	log.Println("Runnin on :7777")
	log.Fatal(http.ListenAndServe(":7777", router))
}

func Metrics(next http.Handler) http.Handler {
	totalRequest := expvar.NewInt("total_request_received")
	totalResponse := expvar.NewInt("total_response_sent")
	totalProcessingTime := expvar.NewInt("total_processing_time_us")
	totalreponsesentbystatus := expvar.NewMap("total_responses_sent_by_status")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		totalRequest.Add(1)
		next.ServeHTTP(w, r)
		metrics := httpsnoop.CaptureMetrics(next, w, r)
		totalResponse.Add(1)

		duration := time.Since(start).Microseconds()
		totalProcessingTime.Add(duration)
		totalreponsesentbystatus.Add(strconv.Itoa(metrics.Code), 1)
	})
}
