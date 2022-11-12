package main

import (
	"app/app"
	"app/app/auth"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Warnf("failed to load .env file: %v", err)
	}

	port := flag.String("p", "8080", "port to serve on")
	directory := flag.String("d", "./docs", "the directory of static file to host")
	flag.Parse()

	f, err := os.Stat(*directory)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("%s does not exist", *directory)
		} else {
			log.Fatalf("unable to access %s", *directory)
		}
	}
	if !f.IsDir() {
		log.Fatalf("%s is not a valid directory", *directory)
	}

	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := app.NewRouter(*directory, authenticator)

	log.Infof("Serving %s on HTTP port: %s\n", *directory, *port)

	addr := fmt.Sprintf(":%s", *port)

	if err := http.ListenAndServe(addr, rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
