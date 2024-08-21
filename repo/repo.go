package repo

import (
	"conn/config"
	"conn/models"
	"conn/pkg"
	"database/sql"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type NewServerr struct {
	db *sql.DB
}

func RepoConn() *NewServerr {
	cfg := config.Load(".env")
	db, err := pkg.ConnectToDB(cfg)

	if err != nil {
		log.Println("Xatolik bor connectdb .envdan uqishda...", err)
		return nil
	}

	return &NewServerr{db: db}
}

func (c *NewServerr) CreateBookk(book *models.Books) (string, error) {
	query := `insert into books(title, author, published_date, isbn) values($1, $2, $3, $4)`
	_, err := c.db.Exec(query, book.Title, book.Author, time.Now().Format(time.RFC3339), book.Isbn)

	if err != nil {
		log.Println("Error on exec", err)
		return "xatolik bor database saqalaganda", err
	}

	resp := "Database saqlandi"
	return resp, nil
}

func (c *NewServerr) GetBook(bookID string) (*models.Books, error) {
	id, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Xatolik atoida", err)
		return nil, err
	}
	query := `select id, title, author, published_date, isbn from books where id=$1`
	book := &models.Books{}

	err = c.db.QueryRow(query, id).Scan(&book.Id, &book.Title, &book.Author, &book.PublishedDate, &book.Isbn)
	if err != nil {
		log.Println("Error fetching book:", err)
		return nil, err
	}

	return book, nil
}

func (c *NewServerr) DeleteBook(id string) (string, error) {
	idd, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Xatolik atoida", err)
		return "Xatolik bor", err
	}

	query := `delete from books where id=$1`
	_, err = c.db.Exec(query, idd)
	if err != nil {
		log.Println("Xatolik bor deleteda", err)
		return "s", nil
	}

	return "Deeleted successfully", nil
}

func (c *NewServerr) UpdateBook(id string, book *models.Books) (string, error) {
	idd, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error converting ID:", err)
		return "Error in ID conversion", err
	}

	query := `UPDATE books SET title=$1, author=$2, published_date=$3, isbn=$4 WHERE id=$5`
	_, err = c.db.Exec(query, book.Title, book.Author, book.PublishedDate, book.Isbn, idd)
	if err != nil {
		log.Println("Error updating book:", err)
		return "Error updating book", err
	}

	return "Book updated successfully", nil
}
