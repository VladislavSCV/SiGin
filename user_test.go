package main

import (
	"bytes"
	"log"
	"io"
	"net/http"
	"testing"
)

func TestPingPong(t *testing.T) {
	var bodyBytes []byte
	var err error
	var conn *http.Response
	if conn, err = http.Get("http://127.0.0.1:8080/ping"); err != nil {
		t.Fatalf("http.Get: %v", err)
	}
	defer conn.Body.Close()
	bodyBytes, err = io.ReadAll(conn.Body)
	if err != nil {
		t.Fatalf("io.ReadAll: %v", err)
	}
	if !bytes.Equal(bodyBytes, []byte(`{"message":"pong"}`)) {
		t.Errorf("got %q, want %q", string(bodyBytes), `{"message":"pong"}`)
	}
}

func TestUserInList(t *testing.T) {
	conn, err := http.Get("http://127.0.0.1:8080/getUsers")
	if err != nil {
		log.Print("OK")
	}
	if conn.Body == nil {
		log.Print("Maybe useList is empty")
	} 
	if conn.StatusCode != http.StatusOK {
		log.Fatalf("Http req is bad")
	}
}

