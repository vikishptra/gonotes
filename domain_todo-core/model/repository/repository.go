package repository

import (
	"context"

	"vikishptra/domain_todo-core/model/entity"
	"vikishptra/domain_todo-core/model/vo"
)

type SaveTodoRepo interface {
	SaveTodo(ctx context.Context, obj *entity.Todo) error
}

type FindOneCheckedRepo interface {
	FindOneChecked(ctx context.Context, todoID vo.TodoID) (*entity.Todo, error)
}

type GetAllTodoRepo interface {
	GetAllTodo(ctx context.Context) ([]*entity.Todo, int64, error)
}
