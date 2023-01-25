package service

import (
	"github.com/dinachka/todo-app"
	"github.com/dinachka/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) CreateList(userId int, list todo.TodoList) (int, error) {
	return s.repo.CreateList(userId, list)
}

func (s *TodoListService) GetLists(userId int) ([]todo.TodoList, error) {
	return s.repo.GetLists(userId)
}

func (s *TodoListService) GetListById(userId, listId int) (todo.TodoList, error) {
	return s.repo.GetListById(userId, listId)

}
