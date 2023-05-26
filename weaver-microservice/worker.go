package main

import (
	"context"

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
    result := wo
    result.id = result.id + 1

    return result, nil
}
