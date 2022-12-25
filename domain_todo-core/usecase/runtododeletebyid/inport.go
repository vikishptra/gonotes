package runtododeletebyid

import (
	"vikishptra/domain_todo-core/model/vo"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	TodoID vo.TodoID
}

type InportResponse struct {
	Pesan string `json:"message"`
}
