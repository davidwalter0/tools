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

	// Assume if configuration loads from file, that the environment and
	// flags don't need to be processed
	// err = cfg.Parse(app.WebAuthn)
	// if err != nil || !app.valid() {
	// 	log.Info(err)
	// }
	// fmt.Println("------------------------------------------------------------------------")
	// app.DumpJSON()
	// err = app.Parse()
	// fmt.Println("------------------------------------------------------------------------")
	// app.DumpJSON()
	// if err != nil || !app.valid() {
	// 	log.Info(err)
	// }
	// if app.WebAuthn.Debug {
	// 	fmt.Println("------------------------------------------------------------------------")
	// 	app.DumpYAML()
	// 	fmt.Println("------------------------------------------------------------------------")
	// 	app.DumpJSON()
	// }

	// if err = cfg.Parse(app.WebAuthn); err != nil {
	// 	log.Info(err)
	// 	os.Exit(1)
	// }

	// if err = app.Parse(); err != nil {
	// 	log.Info(err)
	// 	os.Exit(1)
	// }

	if err = cfg.Parse(app.WebAuthn); err != nil {
		log.Info(err)
		os.Exit(1)
	}

	if err = app.Parse(); err != nil {
		log.Info(err)
		os.Exit(1)
	}

	fmt.Println(app.Parse())

	// if app.WebAuthn.Debug {
	fmt.Println("------------------------------------------------------------------------")
	app.DumpYAML()
	fmt.Println("------------------------------------------------------------------------")
	app.DumpJSON()
	// }

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
