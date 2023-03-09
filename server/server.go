package server

import (
	"context"
	"github.com/gorilla/mux"
	"golang_test_task/config"
	"golang_test_task/handler/create"
	"golang_test_task/handler/get"
	"golang_test_task/middleware"
	"log"
	"net/http"
	"time"
)

func InitServer(ctx context.Context) {
	conf := config.GetConfig()
	server := &http.Server{
		Addr:         ":" + conf.ServerPort,
		Handler:      getRoutes(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	<-ctx.Done()
	err := server.Shutdown(context.Background())
	if err != nil {
		log.Println("failed to shutdown the server gracefully: " + err.Error())
	}
}

func getRoutes() *mux.Router {
	m := mux.NewRouter()
	m.Use(middleware.JSON)
	m.HandleFunc("/task", create.Handler).Methods(http.MethodPost)
	m.HandleFunc("/task/{id}", get.Handle).Methods(http.MethodGet)

	return m
}
