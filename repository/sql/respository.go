package sql

import (
	"log"
	"strconv"

	"database/sql"

	"github.com/pkg/errors"
	"github.com/sirdeggen/familiarisation/shortener"

	_ "github.com/lib/pq"
)

type sqlRepository struct {
	client *sql.DB
}

func newSqlClient(sqlURL string) (*sql.DB, error) {
	client, err := sql.Open("postgres", sqlURL)
	if err != nil {
		return nil, err
	}
	err2 := client.Ping()
	if err2 != nil {
		return nil, err2
	}
	return client, nil
}

func NewSqlRepository(sqlURL string) (shortener.RedirectRepository, error) {
	repo := &sqlRepository{}
	client, err := newSqlClient(sqlURL)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewSqlRepository")
	}
	repo.client = client
	return repo, nil
}

func (r *sqlRepository) Find(code string) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	sqlStatement := `SELECT * FROM redirects WHERE code=$1`
	row := r.client.QueryRow(sqlStatement, code)
	var url, c, creAt string
	row.Scan(&url, &c, &creAt)
	if len(url) == 0 {
		return nil, errors.Wrap(shortener.ErrRedirectNotFound, "repository.Redirect.Find")
	}
	createdAt, err := strconv.ParseInt(creAt, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.Find")
	}
	redirect.Code = c
	redirect.URL = url
	redirect.CreatedAt = createdAt
	return redirect, nil
}

func (r *sqlRepository) Store(redirect *shortener.Redirect) error {
	data := map[string]interface{}{
		"code":       redirect.Code,
		"url":        redirect.URL,
		"created_at": redirect.CreatedAt,
	}
	sqlStatement := `
	INSERT INTO redirects (url, code, created_at)
	VALUES ($1, $2, $3)`
	_, err := r.client.Exec(sqlStatement, data["url"], data["code"], data["created_at"])
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		return errors.Wrap(err, "repository.Redirect.Store")
	}
	return nil
}
