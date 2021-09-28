package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	h "github.com/sirdeggen/familiarisation/api"
	mr "github.com/sirdeggen/familiarisation/repository/mongodb"
	rr "github.com/sirdeggen/familiarisation/repository/redis"
	sr "github.com/sirdeggen/familiarisation/repository/sql"
	"github.com/sirdeggen/familiarisation/shortener"
)

// repo <- service -> serialiser -> http
func main() {
	repo := chooseRepo()
	service := shortener.NewRedirectService(repo)
	handler := h.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{code}", handler.Get)
	r.Post("/", handler.Post)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :8000")
		errs <- http.ListenAndServe(httpPort(), r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)

}

func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func chooseRepo() shortener.RedirectRepository {
	repoChoice := "sql" //os.Getenv("URL_DB")
	log.Println("chooseRepo: ", repoChoice)
	switch repoChoice {
	case "redis":
		redisURL := "redis://localhost:6379" //os.Getenv("REDIS_URL")
		repo, err := rr.NewRedisRepository(redisURL)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case "mongo":
		mongoURL := "mongodb://localhost/shortner" //os.Getenv("MONGO_URL")
		mongodb := "shortner"                      //os.Getenv("MONGO_DB")
		mongoTimeout, _ := strconv.Atoi("30")      //os.Getenv("MONGO_TIMEOUT"))
		repo, err := mr.NewMongoRepository(mongoURL, mongodb, mongoTimeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case "sql":
		sqlURL := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			"localhost", 5432, "postgres", "greatPassword1!", "postgres") //os.Getenv("MONGO_URL")
		repo, err := sr.NewSqlRepository(sqlURL)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
