package getalltodo

import (
	"context"

	"vikishptra/domain_todo-core/model/errorenum"
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

	todoObjs, count, page, err := r.outport.GetAllTodo(ctx, req.Page, req.Size)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(todoObjs)
	if res.Items == nil {
		return nil, errorenum.DataNull
	}
	res.Page = page

	return res, nil
}
