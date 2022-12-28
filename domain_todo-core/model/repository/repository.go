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
	GetAllTodoByPagination(ctx context.Context, page, size int) ([]*entity.Todo, int64, int, error)
}

type DeleteOneTodoByIDRepo interface {
	DeleteOneTodoByID(ctx context.Context, todoID vo.TodoID) (*entity.Todo, error)
}

type FindMessageTodoEmptyRepo interface {
	FindMessageTodoEmpty(ctx context.Context, obj *entity.Todo) error
}

type GetTodoByIDRepo interface {
	GetTodoByID(ctx context.Context, someID vo.TodoID) ([]*entity.Todo, error)
}
