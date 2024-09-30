package service

import (
	"database/sql"
	"errors"
	"go-todo/models"
	"go-todo/repository"
)

func (s *Service) CreateTodo(todo models.Todo) (models.Todo, error) {
	newTodo, err := repository.CreateTodo(s.db, todo)
	if err != nil {
		return models.Todo{}, errors.New("at Service CreateTodo")
	}

	return newTodo, nil
}

func (s *Service) ReadTodos(page int) ([]models.Todo, error) {
	todoList, err := repository.ReadTodos(s.db, page)
	if err != nil {
		return nil, errors.New("at Service ReadTodos")
	}

	if len(todoList) == 0 {
		return nil, errors.New("at Service ReadTodos, no data found")
	}

	return todoList, nil
}

func (s *Service) ReadTodoByID(id int) (models.Todo, error) {
	foundTodo, err := repository.ReadTodoByID(s.db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Todo{}, errors.New("at Service ReadTodoByID, no data found")
		}

		return models.Todo{}, errors.New("at Service ReadTodoByID")
	}

	return foundTodo, nil
}

func (s *Service) CompleteTodo(id int) (models.Todo, error) {
	_, err := repository.ReadTodoByID(s.db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Todo{}, errors.New("at Service CompleteTodo, no data found")
		}

		return models.Todo{}, errors.New("at Service CompleteTodo(ReadTodoByID)")
	}

	completedTodo, err := repository.CompleteTodo(s.db, id)
	if err != nil {
		return models.Todo{}, errors.New("at Service CompleteTodo")
	}
	return completedTodo, nil
}

func (s *Service) DeleteTodo(id int) (models.Todo, error) {
	_, err := repository.ReadTodoByID(s.db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Todo{}, errors.New("at Service DeleteTodo, no data found")
		}

		return models.Todo{}, errors.New("at Service DeleteTodo(ReadTodoByID)")
	}

	deletedTodo, err := repository.DeleteTodo(s.db, id)
	if err != nil {
		return models.Todo{}, errors.New("at Service DeleteTodo")
	}
	return deletedTodo, nil
}
