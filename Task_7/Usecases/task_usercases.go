package usecases

import (
	domain "Task_7/Domain"
	"context"
	"time"
)

type taskUsecase struct {
	taskRepository domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(taskRepo domain.TaskRepository, timeout time.Duration) domain.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepo,
		contextTimeout: timeout,
	}
}

func (tu *taskUsecase) CreateTask(c context.Context, task *domain.Task) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.CreateTask(ctx, task)
}

func (tu *taskUsecase) GetTaskByID(c context.Context, taskID string) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.GetTaskByID(ctx, taskID)
}

func (tu *taskUsecase) GetAllTasks(c context.Context) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.GetAllTasks(ctx)
}

func (tu *taskUsecase) UpdateTask(c context.Context, taskID string, task domain.Task) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.UpdateTask(ctx, taskID, task)
}

func (tu *taskUsecase) DeleteTask(c context.Context, taskID string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.DeleteTask(ctx, taskID)
}
