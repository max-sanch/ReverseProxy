package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/max-sanch/ReverseProxy/pkg/handler"
	"github.com/max-sanch/ReverseProxy/pkg/service"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "rdb:6379",
		Password: "",
		DB:       0,
	})

	defer rdb.Close()

	services := service.NewService(rdb)
	handlers := handler.NewHandler(services)

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handlers.WriteResponse))

	srv := &http.Server{
		Addr:    ":" + viper.GetString("port"),
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}