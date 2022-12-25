package runtodochecked

import "vikishptra/domain_todo-core/model/repository"

type Outport interface {
	repository.FindOneCheckedRepo
	repository.SaveTodoRepo
}
