package main

import (
	"github.com/max-sanch/ReverseProxy/pkg/handler"
	"github.com/max-sanch/ReverseProxy/pkg/service"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing config: %s", err.Error())
	}

	services := service.NewService()
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