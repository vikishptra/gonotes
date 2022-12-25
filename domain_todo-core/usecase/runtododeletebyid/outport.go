package runtododeletebyid

import "vikishptra/domain_todo-core/model/repository"

type Outport interface {
	repository.DeleteOneTodoByIDRepo
	repository.FindOneCheckedRepo
}
