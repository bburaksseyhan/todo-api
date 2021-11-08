package db

import (
	"github.com/bburaksseyhan/todo-api/src/pkg/model"

	"github.com/gocql/gocql"
)

type TodoRepository interface {
	Save(todo model.Todo) (*model.Todo, error)
	GetById(id string) (*model.Todo, error)
}

type todoRepository struct {
	session *gocql.Session
}

func NewTodoRepository(session *gocql.Session) TodoRepository {
	return &todoRepository{session: session}
}

func (t *todoRepository) Save(todo model.Todo) (*model.Todo, error) {

	var query string = "INSERT INTO todo(id,title,content) VALUES(?,?,?)"

	if err := t.session.Query(query, todo.Id, todo.Title, todo.Content).Exec(); err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *todoRepository) GetById(id string) (*model.Todo, error) {

	var todo model.Todo

	var query string = `SELECT id,title,content FROM todo where id=?`

	if err := t.session.Query(query, id).Scan(&todo.Id, &todo.Title, &todo.Content); err != nil {

		if err == gocql.ErrNotFound {
			return nil, err
		}

		return nil, err
	}

	return &todo, nil
}
