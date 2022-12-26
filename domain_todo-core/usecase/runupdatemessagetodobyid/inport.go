package runupdatemessagetodobyid

import (
	"vikishptra/domain_todo-core/model/entity"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.TodoUpdateRequest
}

type InportResponse struct {
	Todo *entity.Todo
}
