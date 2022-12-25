package runtodochecked

import (
	"vikishptra/domain_todo-core/model/entity"
	"vikishptra/domain_todo-core/model/vo"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	TodoId vo.TodoID
}

type InportResponse struct {
	Todo *entity.Todo
}
