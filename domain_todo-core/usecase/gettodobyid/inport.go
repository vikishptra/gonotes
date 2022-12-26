package gettodobyid

import (
	"vikishptra/domain_todo-core/model/vo"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	TodoID vo.TodoID
}

type InportResponse struct {
	Items []any `json:"items"`
}
