package main

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

type Worker interface {
	Work(context.Context, WorkObject) (WorkObject, error)
	Work2(context.Context, WorkObject2) (WorkObject2, error)
}

type worker struct {
	weaver.Implements[Worker]
}

type WorkObject struct {
	weaver.AutoMarshal
	id      int
	action  string
	message string
}

func (w *worker) Work(_ context.Context, wo WorkObject) (WorkObject, error) {
	result := wo
	result.id = result.id + 1

	return result, nil
}

func (w *worker) Work2(_ context.Context, wo WorkObject2) (WorkObject2, error) {
	result := wo
	result.id = result.id + 1

	return result, nil
}

type WorkObject2 struct {
	weaver.AutoMarshal
	id        int
	action    string
	message1  string
	message2  string
	message3  string
	message4  string
	message5  string
	message6  string
	message7  string
	message8  string
	message9  string
	message10 string
	message11 string
	message12 string
	message13 string
	message14 string
	message15 string
	message16 string
	message17 string
	message18 string
	message19 string
	message20 string
}
