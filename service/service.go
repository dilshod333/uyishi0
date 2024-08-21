package service

import (
	"conn/models"
	"conn/repo"
	"encoding/json"
	"log"
	"net/http"
)

type Server struct {
	ss *repo.NewServerr
}

func NewServer() (*Server, error) {
	r := repo.RepoConn()

	return &Server{ss: r}, nil
}

func (s *Server) CreateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Wrong", http.StatusMethodNotAllowed)
		return
	}
	log.Println("Createbook servicega keldi...>>>>>>>")
	w.Header().Set("Content-Type", "application/json")
	var book models.Books
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Error on decoding", 500)
		return
	}
	resp, err := s.ss.CreateBookk(&book)

	if err != nil {
		http.Error(w, "Error on saving database", 500)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (s *Server) GetBookID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Error on gettting id", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	idStr := r.URL.Path[len("/books/"):]
	log.Println("ID>>>>>>>>>>>>", idStr)
	resp, err := s.ss.GetBook(idStr)
	if err != nil {
		http.Error(w, "Error on getting book", 500)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (s *Server) DeleteBookById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	idstr := r.URL.Path[len("/books/"):]
	resp, err := s.ss.DeleteBook(idstr)
	if err != nil {
		http.Error(w, "Error on delete book  service", 500)
		return
	}
	json.NewEncoder(w).Encode(resp)

}


func (s *Server) UpdateBookk(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) 
		return 
	}

	idstr := r.URL.Path[len("/books/"):]
	var b models.Books
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "Decode error", http.StatusConflict)
		return 
	}

	resp, err := s.ss.UpdateBook(idstr, &b)
	if err != nil {
		http.Error(w, "Error on update on repo", 500)
		return 
	}

	json.NewEncoder(w).Encode(resp)


}