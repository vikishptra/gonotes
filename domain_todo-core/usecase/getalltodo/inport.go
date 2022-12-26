package getalltodo

import (
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	Page int
	Size int
}

type InportResponse struct {
	Count int64 `json:"count"`
	Items []any `json:"items"`
	Page  int   `json:"page"`
}
