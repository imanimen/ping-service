package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	urls := strings.Split(os.Getenv("PING_URLS"), ",")
	for _, url := range urls {
		go pingURL(url)
	}

	// create a channel for listening to the signal

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// wait for the stop channel to receive a signal
	
	<-stop

	log.Println("Shutting down")

}


func pingURL(url string) {
	url = strings.TrimSpace(url)
	for {
		_, err := http.Get(url)
		log.Println("Pinging " + url)
		if err != nil {
			log.Println("Error: ", err.Error())
		}
		time.Sleep(time.Second * 5)
	}
}