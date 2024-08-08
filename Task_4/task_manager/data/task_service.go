package data

import (
	"errors"
	"sync"
	"task_manager/models"
)

type TaskService struct {
	//to ensure thread safe access to the tasks map
	//for concurency control
	mu sync.Mutex

	tasks map[string]models.Task
}

// return the pointer to Taskservice struct
func NewTaskService() *TaskService {

	//return the pointer to newly created TaskService instance
	return &TaskService{
		tasks: make(map[string]models.Task),
	}
}

func (s *TaskService) GetAllTasks() []models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	//creating a list to store the task
	taskList := make([]models.Task, 0, len(s.tasks))

	for _, task := range s.tasks {
		taskList = append(taskList, task)
	}
	return taskList
}

func (s *TaskService) GetTaskById(id string) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, ok := s.tasks[id]
	if !ok {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

func (s *TaskService) AddTask(task models.Task) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tasks[task.ID] = task
}

func (s *TaskService) UpdateTask(id string, UpdatedTask models.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, ok := s.tasks[id]
	if !ok {
		return errors.New("task not found")

	}
	if UpdatedTask.Title != "" {
		task.Title = UpdatedTask.Title
	}
	if UpdatedTask.Description != "" {
		task.Description = UpdatedTask.Description
	}
	if UpdatedTask.DueDate.IsZero() {
		task.DueDate = UpdatedTask.DueDate
	}
	if UpdatedTask.Status != "" {
		task.Status = UpdatedTask.Status
	}

	s.tasks[id] = task
	return nil

}

func (s *TaskService) DeleteTask(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.tasks[id]
	if !ok {
		return errors.New("task not found")
	}
	delete(s.tasks, id)
	return nil
}
