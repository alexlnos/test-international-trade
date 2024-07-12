package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"test-international-trade/config"
	_ "test-international-trade/docs"
	"test-international-trade/internal/handler"
	"test-international-trade/internal/repository"
	"test-international-trade/internal/usecase"
)

// @title Encryption Service API
// @version 1.0
// @description This is a service for encrypting strings using MD5 and SHA256 algorithms
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: config.RedisAddr,
	})

	repo := repository.NewRedisRepository(redisClient)

	useCase := usecase.NewEncryptUseCase(repo)
	httpHandler := handler.NewHTTPHandler(useCase)

	r := mux.NewRouter()
	r.HandleFunc("/encrypt", httpHandler.Encrypt).Methods("POST")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	handler := cors.Default().Handler(r)
	log.Printf("Server starting on port %s", config.ServerPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), handler); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
