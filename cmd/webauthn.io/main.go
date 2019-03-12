package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	cfg "github.com/davidwalter0/go-cfg"
	"github.com/davidwalter0/tools/x/webauthn.io/config"
	log "github.com/davidwalter0/tools/x/webauthn.io/logger"
	"github.com/davidwalter0/tools/x/webauthn.io/models"
	"github.com/davidwalter0/tools/x/webauthn.io/server"
)

var app = &App{
	WebAuthn: &config.Config{},
}

func main() {
	var err error
	// config, err := config.LoadConfig("config.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// app.WebAuthn = config

	app.DumpYAML()
	app.DumpJSON()

	if err = cfg.Parse(app.WebAuthn); err != nil {
		log.Info(err)
		os.Exit(1)
	}

	if err = app.Parse(); err != nil {
		log.Info(err)
		os.Exit(1)
	}

	fmt.Println(app.Parse())

	app.DumpYAML()
	app.DumpJSON()

	err = models.Setup(app.WebAuthn)
	if err != nil {
		log.Fatal(err)
	}

	err = log.Setup(app.WebAuthn)
	if err != nil {
		log.Fatal(err)
	}

	server, err := server.NewServer(app.WebAuthn)
	if err != nil {
		log.Fatal(err)
	}
	go server.Start()

	// Handle graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGINT)

	<-c
	log.Info("Shutting down...")
	server.Shutdown()
}
