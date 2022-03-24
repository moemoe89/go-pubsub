package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cloud.google.com/go/pubsub"
)

var (
	//projectID of GCP project.
	projectID = "go-app-275713"
	// topicID of PubSub topic.
	topicID = "medium"
)

// publishHandler handler for publish message.
func publishHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := req.Context()

	// Init PubSub Client.
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		msg := "failed to init pubsub client"
		log.Printf(msg+": %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": msg,
		})

		return
	}
	defer client.Close()

	// Create a references to a topic.
	t := client.Topic(topicID)
	// Publish the message to the topic.
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte("Hello world!"),
		Attributes: map[string]string{
			"origin":   "golang",
			"username": "gcp",
		},
	})

	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		msg := "failed to init publish message"
		log.Printf(msg+": %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": msg,
		})

		return
	}

	// Write response on JSON.
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": fmt.Sprintf("published message with custom attributes; msg id: %s", id),
	})
}

func main() {
	// Init HTTP server
	mux := http.NewServeMux()
	mux.HandleFunc("/publish", publishHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start running the server.
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve: %v\n", err)
		}
	}()
	log.Printf("server is starting at %s...", srv.Addr)

	// Receive signal to shutdown the server.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("signal %d received, shutting down gracefully...", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	// Gracefully shutdown the server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("could not gracefully shutdown the server: %v\n", err)
	}
	log.Println("finished graceful shutdown")
}
