package gettodobyid

import (
	"context"

	"vikishptra/domain_todo-core/model/errorenum"
	"vikishptra/shared/util"
)

type getTodoByIDInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &getTodoByIDInteractor{
		outport: outputPort,
	}
}

func (r *getTodoByIDInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...

	todoObjs, err := r.outport.GetTodoByID(ctx, req.TodoID)
	if err != nil {
		return nil, err
	}
	res.Items = util.ToSliceAny(todoObjs)
	if res.Items == nil {
		return nil, errorenum.DataNull
	}
	return res, nil
}
