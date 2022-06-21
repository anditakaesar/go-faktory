package main

import (
	"fmt"
	"time"

	faktory "github.com/contribsys/faktory/client"
)

func main() {
	fmt.Println("Pushing my first job")
	server := faktory.Server{
		Network:  "tcp",
		Address:  "localhost:7419",
		Username: "faktory",
		Password: "andita2faktory",
		Timeout:  1 * time.Second,
		TLS:      nil,
	}

	client, err := server.Open()
	if err != nil {
		panic(err)
	}

	job := faktory.NewJob("SomeJob", "hello world")
	err = client.Push(job)

	if err != nil {
		panic(err)
	}

	fmt.Println("Job pushed")
}
