package main

import (
	"context"
	"fmt"

	worker "github.com/contribsys/faktory_worker_go"
)

func main() {
	mgr := worker.NewManager()
	// register someJobWorker func to execute the job of type somejob
	mgr.Register("SomeJob", someJobWorker)
	mgr.Register("SomeJobAgain", someJobWorkerAgain)
	mgr.Run()
}

func someJobWorker(ctx context.Context, args ...interface{}) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occured:", err)
		}
	}()

	message := args[0].(int)
	print("Job Executed1 message: ", message)
	return nil
}

func someJobWorkerAgain(ctx context.Context, args ...interface{}) error {
	message := args[0].(string)
	print("Job Executed message: ", message)
	return nil
}
