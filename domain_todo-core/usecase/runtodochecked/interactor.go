package runtodochecked

import (
	"context"
)

type runTodoCheckedInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runTodoCheckedInteractor{
		outport: outputPort,
	}
}

func (r *runTodoCheckedInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	todoObj, err := r.outport.FindOneChecked(ctx, req.TodoId)
	if err != nil {
		return nil, err
	}
	todoObj.SetTrue()
	r.outport.SaveTodo(ctx, todoObj)
	res.Todo = todoObj

	return res, nil
}
