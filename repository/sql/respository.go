package redis

import (
	"fmt"
	"log"
	"strconv"

	"github.com/pkg/errors"
	"github.com/sirdeggen/familiarisation/shortener"
)

type sqlRepository struct {
	client *sql.Client
}

func newSqlClient(sqlURL string) (*sql.Client, error) {
	log.Println("newSqlClient")
	opts, err := sql.ParseURL(sqlURL)
	if err != nil {
		return nil, err
	}
	client := sql.NewClient(opts)
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewSqlRepository(sqlURL string) (shortener.RedirectRepository, error) {
	log.Println("NewRedisRepository")
	repo := &sqlRepository{}
	client, err := newSqlClient(sqlURL)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewSqlRepository")
	}
	repo.client = client
	return repo, nil
}

func (r *sqlRepository) generateKey(code string) string {
	log.Println("generateKey")
	return fmt.Sprintf("redirect:%s", code)
}

func (r *sqlRepository) Find(code string) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	key := r.generateKey(code)
	data, err := r.client.HGetAll(key).Result()
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.Find")
	}
	if len(data) == 0 {
		return nil, errors.Wrap(shortener.ErrRedirectNotFound, "repository.Redirect.Find")
	}
	createdAt, err := strconv.ParseInt(data["created_at"], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.Find")
	}
	redirect.Code = data["code"]
	redirect.URL = data["url"]
	redirect.CreatedAt = createdAt
	return redirect, nil
}

func (r *sqlRepository) Store(redirect *shortener.Redirect) error {
	log.Println("40")
	key := r.generateKey(redirect.Code)
	log.Println("41")
	data := map[string]interface{}{
		"code":       redirect.Code,
		"url":        redirect.URL,
		"created_at": redirect.CreatedAt,
	}
	log.Println("42")
	_, err := r.client.HMSet(key, data).Result()
	log.Println("43")
	if err != nil {
		log.Println("44")
		return errors.Wrap(err, "repository.Redirect.Store")
	}
	return nil
}
