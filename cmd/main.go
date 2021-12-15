package main

import (
	"context"
	"errors"
	"github.com/inspectorvitya/x-technology-test/internal/application"
	"github.com/inspectorvitya/x-technology-test/internal/config"
	httpserver "github.com/inspectorvitya/x-technology-test/internal/server"
	"github.com/inspectorvitya/x-technology-test/internal/storage/pgsql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(cfg)
	}
	db, err := pgsql.New(cfg)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	ctx := context.Background()

	app := application.New(db)
	app.Init(ctx)
	server := httpserver.New(cfg, app)
	go func() {
		log.Println("http server start...")
		if err := server.Start(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Println("http server http stopped....")
			} else {
				log.Fatalln(err)
			}
		}
	}()
	<-stop
	ctxClose, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = server.Stop(ctxClose)
	if err != nil {
		log.Fatalln(err)
	}
}
