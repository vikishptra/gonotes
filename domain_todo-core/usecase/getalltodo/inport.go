package getalltodo

import (
	"vikishptra/domain_todo-core/model/entity"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.Todo
}

type InportResponse struct {
	Count int64 `json:"count"`
	Items []any `json:"items"`
}
