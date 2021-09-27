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
	"github.com/sirdeggen/familiarisation/shortener"
)

// repo <- service -> serialiser -> http
func main() {
	log.Println("1")
	repo := chooseRepo()
	log.Println("2")
	service := shortener.NewRedirectService(repo)
	log.Println("3")
	handler := h.NewHandler(service)
	log.Println("4")

	r := chi.NewRouter()
	log.Println("5")
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{code}", handler.Get)
	log.Println("6")
	r.Post("/", handler.Post)
	log.Println("7")

	errs := make(chan error, 2)
	log.Println("8")
	go func() {
		fmt.Println("Listening on port :8000")
		log.Println("9")
		errs <- http.ListenAndServe(httpPort(), r)
		log.Println("10")
	}()

	go func() {
		log.Println("11")
		c := make(chan os.Signal, 1)
		log.Println("12")
		signal.Notify(c, syscall.SIGINT)
		log.Println("13")
		errs <- fmt.Errorf("%s", <-c)
		log.Println("14")
	}()

	fmt.Printf("Terminated %s", <-errs)

}

func httpPort() string {
	log.Println("httpPort")
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func chooseRepo() shortener.RedirectRepository {
	repoChoice := "mongo" //os.Getenv("URL_DB")
	log.Println("chooseRepoooo", repoChoice)
	switch repoChoice {
	case "redis":
		redisURL := "redis://localhost:6379" //os.Getenv("REDIS_URL")
		repo, err := rr.NewRedisRepository(redisURL)
		if err != nil {
			log.Println("BAD REDIS")
			log.Fatal(err)
		}
		log.Println("GOOD REDIS")
		return repo
	case "mongo":
		log.Println("chooseRepo")
		mongoURL := "mongodb://localhost/shortner" //os.Getenv("MONGO_URL")
		mongodb := "shortner"                      //os.Getenv("MONGO_DB")
		mongoTimeout, _ := strconv.Atoi("30")      //os.Getenv("MONGO_TIMEOUT"))
		repo, err := mr.NewMongoRepository(mongoURL, mongodb, mongoTimeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
