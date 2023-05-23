package main

import (
	"context"
	"log"

	"github.com/ServiceWeaver/weaver"
)

type Worker interface {
    Work(context.Context, WorkObject) (WorkObject, error)
}

type WorkObject struct {
    weaver.AutoMarshal
    id int
    action string
    message string
}

type worker struct {
    weaver.Implements[Worker]
}

func (w *worker) Work(_ context.Context, wo WorkObject) (WorkObject, error) {
    log.Println("worker: received request")

    result := wo
    result.id = result.id + 1

    return result, nil
}
