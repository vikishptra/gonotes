package runtodocreate

import "vikishptra/domain_todo-core/model/repository"

type Outport interface {
	repository.SaveTodoRepo
	repository.FindMessageTodoEmptyRepo
}
