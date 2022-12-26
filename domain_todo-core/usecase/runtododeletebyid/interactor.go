package runtododeletebyid

import (
	"context"

	"vikishptra/domain_todo-core/model/errorenum"
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
	if !todoObj.Checked {
		return nil, errorenum.SomethingError
	}
	r.outport.DeleteOneTodoByID(ctx, todoObj.ID)
	res.Pesan = "ok success"

	return res, nil
}
