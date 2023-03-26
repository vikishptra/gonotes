package entity

import (
	"strings"
	"time"

	"vikishptra/domain_todo-core/model/errorenum"
	"vikishptra/domain_todo-core/model/vo"
)

type Todo struct {
	ID          vo.TodoID `bson:"_id" json:"id"`
	Created     time.Time `bson:"created" json:"created"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Checked     bool      `json:"checked"`
}

type TodoCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
}

func (r TodoCreateRequest) Validate() error {

	if strings.TrimSpace(r.Title) == "" || strings.TrimSpace(r.Description) == "" {
		return errorenum.MessageEmpty
	}

	return nil
}

func NewTodo(req TodoCreateRequest) (*Todo, error) {

	id, err := vo.NewTodoID(req.RandomString, req.Now, req.Title, req.Description)
	if err != nil {
		return nil, err
	}

	var obj Todo
	obj.ID = id
	obj.Created = req.Now
	obj.Title = req.Title
	obj.Description = req.Description

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

func (r TodoUpdateRequest) Validate() error {

	if strings.TrimSpace(r.Title) == "" || strings.TrimSpace(r.Description) == "" {
		return errorenum.MessageEmpty
	}

	return nil
}

type TodoUpdateRequest struct {
	ID          vo.TodoID `uri:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

func (r *Todo) Update(req TodoUpdateRequest) error {
	r.Title = req.Title
	r.Description = req.Description

	return req.Validate()
}
