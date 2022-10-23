package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	startServer()

	resp, err := http.Post("http://localhost:3000/hello", "application/json", makeRequest())
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, resp.Body)
}

func makeRequest() io.Reader {
	req := &Request{Name: "Gopher"}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(req)
	return buf
}

func startServer() {
	server := NewServer()
	go server.ListenAndServe()
	time.Sleep(100 * time.Millisecond)
}
