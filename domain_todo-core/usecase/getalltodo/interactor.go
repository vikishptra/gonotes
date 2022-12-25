package getalltodo

import (
	"context"

	"vikishptra/shared/util"
)

type getAllTodoInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &getAllTodoInteractor{
		outport: outputPort,
	}
}

func (r *getAllTodoInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...

	todoObjs, count, err := r.outport.GetAllTodo(ctx)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(todoObjs)

	//!

	return res, nil
}
