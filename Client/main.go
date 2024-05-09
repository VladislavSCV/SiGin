package main

import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"

	_ "time"
)

type Time struct {
	Day   int
	Hour  int
	Minute int
	Seconds int `json:Seconds`
}

var time Time;

func main() {
	for true {
		// fmt.Print("\033[H\033[2J")
		resp, err := http.Get("http://127.0.0.1:8000/time")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		var time Time
		if err := json.NewDecoder(resp.Body).Decode(&time); err != nil {
			log.Fatal(err)
		}

	
		log.Printf("Day: %v Hour: %v Minute: %v Second: %v", time.Day, time.Hour, time.Minute, time.Seconds)
		
	}
}
