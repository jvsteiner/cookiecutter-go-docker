package main

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Greeting string `json:"greeting"`
}
