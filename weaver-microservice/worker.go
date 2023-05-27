package main

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

type Workaholic interface {
	Work(context.Context, WorkObject) (WorkObject, error)
}

type workerholic struct {
	weaver.Implements[Workaholic]
}

type WorkObject struct {
	weaver.AutoMarshal
	id      int
	action  string
	message string
}

func (w *workerholic) Work(_ context.Context, wo WorkObject) (WorkObject, error) {
	result := wo
	result.id = result.id + 1

	return result, nil
}
