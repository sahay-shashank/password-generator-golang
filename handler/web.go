package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/sahay-shashank/password-generator-golang/password"
)

func home(w http.ResponseWriter, r *http.Request) {
	lengthstr := r.URL.Query().Get("length")
	if lengthstr == "" {
		http.Error(w, "Missing Parameter: 'length'", http.StatusBadRequest)
		return
	}
	length, err := strconv.Atoi(lengthstr)
	if err != nil {
		http.Error(w, "Invalid value in parameter 'length'. Provide valid Integer.", http.StatusBadRequest)
		return
	}
	if length < 7 {
		http.Error(w, "Parameter Validation failed: 'length' should be more than 7.", http.StatusBadRequest)
		return
	}
	response := password.Password(length)

	fmt.Fprint(w, response)
}

func Web() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", home)
	log.Printf("Server at http://localhost:" + port)
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	server := &http.Server{Addr: ":" + port}
	go func() {
		log.Fatal(server.ListenAndServe())

	}()

	<-stopChan // wait for SIGINT
	log.Println("Shutting down server...")

	// shut down gracefully, but wait no longer than 5 seconds before halting
	ctx, c := context.WithTimeout(context.Background(), 5*time.Second)
	defer c()
	server.Shutdown(ctx)

	log.Println("Server gracefully stopped")
}
