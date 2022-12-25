package entity

import (
	"strings"
	"time"

	"vikishptra/domain_todo-core/model/errorenum"
	"vikishptra/domain_todo-core/model/vo"
)

type Todo struct {
	ID      vo.TodoID `bson:"_id" json:"id"`
	Created time.Time `bson:"created" json:"created"`
	Message string    `json:"message"`
	Checked bool      `json:"checked"`
}

type TodoCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`
	Message      string    `json:"message"`
}

func (r TodoCreateRequest) Validate() error {

	if strings.TrimSpace(r.Message) == "" {
		return errorenum.MessageEmpty
	}

	return nil
}

func NewTodo(req TodoCreateRequest) (*Todo, error) {

	id, err := vo.NewTodoID(req.RandomString, req.Now, req.Message)
	if err != nil {
		return nil, err
	}

	// add validation and assignment value here ...

	var obj Todo
	obj.ID = id
	obj.Created = req.Now
	obj.Message = req.Message

	return &obj, req.Validate()
}

func (r *Todo) SetTrue() error {

	if r.Checked {
		r.Checked = false
		return nil
	}
	r.Checked = true

	return nil
}

type TodoUpdateRequest struct {
	// add field to update here ...
}

func (r *Todo) Update(req TodoUpdateRequest) error {

	// add validation and assignment value here ...

	return nil
}
