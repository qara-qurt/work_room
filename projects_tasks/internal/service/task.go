package service

import "projects_tasks/internal/repository"

type Task struct {
	repo repository.ITaskRepository
}

func (t Task) CreateTasks() {
	//TODO implement me
	panic("implement me")
}

func (t Task) GetTask() {
	//TODO implement me
	panic("implement me")
}

func (t Task) UpdateTask() {
	//TODO implement me
	panic("implement me")
}

func (t Task) DeleteTask() {
	//TODO implement me
	panic("implement me")
}

func NewTask(repo repository.ITaskRepository) *Task {
	return &Task{
		repo: repo,
	}
}
