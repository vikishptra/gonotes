package runupdatemessagetodobyid

import (
	"context"
)

type runUpdateMessageTodoByIDInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runUpdateMessageTodoByIDInteractor{
		outport: outputPort,
	}
}

func (r *runUpdateMessageTodoByIDInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	todoObj, err := r.outport.FindOneChecked(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if err := todoObj.Update(req.TodoUpdateRequest); err != nil {
		return nil, err
	}
	if err := r.outport.FindMessageTodoEmpty(ctx, todoObj); err != nil {
		return nil, err
	}
	if err := r.outport.SaveTodo(ctx, todoObj); err != nil {
		return nil, err
	}
	res.Todo = todoObj
	return res, nil
}
