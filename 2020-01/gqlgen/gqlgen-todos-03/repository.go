package gqlgen_todos

import (
	"encoding/json"
	"io"
	"os"
	"sync"

	"github.com/friendsofgo/errors"
	"github.com/gofrs/uuid"
)

type TodosRepository struct {
	file *os.File
	data *storedData
	mx   sync.RWMutex
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
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
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
	r.mx.RLock()
	defer r.mx.RUnlock()

	res := make([]*Todo, 0, len(r.data.Todos))
	for _, todo := range r.data.Todos {
		res = append(res, r.mapTodo(todo))
	}

	return res
}

func (r *TodosRepository) mapTodo(todo storedTodo) *Todo {
	return &Todo{
		ID:     todo.ID,
		Text:   todo.Text,
		Done:   todo.Done,
		UserID: todo.UserID,
	}
}

func (r *TodosRepository) FindUserByID(userID string) *User {
	r.mx.RLock()
	defer r.mx.RUnlock()

	user, ok := r.data.Users[userID]
	if !ok {
		return nil
	}
	return r.mapUser(user)
}

func (r *TodosRepository) mapUser(user storedUser) *User {
	return &User{
		ID:   user.ID,
		Name: user.Name,
	}
}

func (r *TodosRepository) InsertTodo(input NewTodo) (*Todo, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	id := uuid.Must(uuid.NewV4()).String()
	t := storedTodo{
		ID:     id,
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	}

	r.data.Todos[id] = t

	_, err := r.file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, errors.Wrap(err, "failed to seek file")
	}
	err = r.file.Truncate(0)
	if err != nil {
		return nil, errors.Wrap(err, "failed to truncate file")
	}

	enc := json.NewEncoder(r.file)
	enc.SetIndent("", "  ")
	err = enc.Encode(r.data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to write data")
	}

	return r.mapTodo(t), nil
}
