package gqlgen_todos

import (
	"encoding/json"
	"os"

	"github.com/friendsofgo/errors"
)

type TodosRepository struct {
	file *os.File
	data *storedData
}

type storedData struct {
	Todos map[string]storedTodo
	Users map[string]storedUser
}

type storedTodo struct {
	ID     string
	Text   string
	Done   bool
	UserID string
}

type storedUser struct {
	ID   string
	Name string
}

func NewTodosRepository(filename string) (*TodosRepository, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "opening repository file")
	}

	data := &storedData{}

	dec := json.NewDecoder(file)
	err = dec.Decode(data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode JSON")
	}

	return &TodosRepository{
		file: file,
		data: data,
	}, nil
}

func (r *TodosRepository) FindAllTodos() []*Todo {
	res := make([]*Todo, 0, len(r.data.Todos))
	for _, todo := range r.data.Todos {
		res = append(res, &Todo{
			ID:   todo.ID,
			Text: todo.Text,
			Done: todo.Done,
			User: r.FindUserByID(todo.UserID),
		})
	}

	return res
}

func (r *TodosRepository) FindUserByID(userID string) *User {
	user, ok := r.data.Users[userID]
	if !ok {
		return nil
	}
	return &User{
		ID:   user.ID,
		Name: user.Name,
	}
}
