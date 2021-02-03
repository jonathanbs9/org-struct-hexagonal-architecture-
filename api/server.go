package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
)

type server struct {
	*http.Server
}

func newServer(listening string, mux *chi.Mux) *Server {
	s := &http.Server{
		Addr:         ":" + listening,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &server{s}
}

func (srv *server) Start() {
	log.Println("Comenzando API cmd")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("No se puede escuchar en : %s |", srv.Addr, err.Error())
		}
	}()
	log.Printf("CMD esta preparado para manejar request! => %s ", srv.Addr)
	srv.gracefulShutDown()
}

func (srv *server) gracefulShutDown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Printf("CMD se está apagando => %s ", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("No se pudo apagar afortunadamente el CMD => %s ", err.Error())
	}
	log.Printf("CMD FRENADO!")
}
