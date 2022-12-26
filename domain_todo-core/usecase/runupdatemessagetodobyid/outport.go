package runupdatemessagetodobyid

import "vikishptra/domain_todo-core/model/repository"

type Outport interface {
	repository.FindOneCheckedRepo
	repository.FindMessageTodoEmptyRepo
	repository.SaveTodoRepo
}
