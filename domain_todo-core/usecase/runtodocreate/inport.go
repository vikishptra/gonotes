package runtodocreate

import (
	"vikishptra/domain_todo-core/gateway/withgorm"
	"vikishptra/domain_todo-core/model/entity"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.TodoCreateRequest
	withgorm.Gateway
}

type InportResponse struct {
	Todo *entity.Todo
}
