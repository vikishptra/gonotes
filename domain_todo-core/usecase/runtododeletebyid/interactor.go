package runtododeletebyid

import (
	"context"
)

type runTodoDeleteByIDInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runTodoDeleteByIDInteractor{
		outport: outputPort,
	}
}

func (r *runTodoDeleteByIDInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	todoObj, err := r.outport.FindOneChecked(ctx, req.TodoID)
	if err != nil {
		return nil, err
	}
	err = r.outport.DeleteOneTodoByID(ctx, todoObj.ID.String())
	if err != nil {
		return nil, err
	}

	return res, nil
}
