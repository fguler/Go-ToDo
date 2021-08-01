package task

import (
	"fmt"
	"time"

	"github.com/fguler/goToDo/domain"
	"github.com/google/uuid"
)

type Service struct {
	repo domain.TaskRepository
}

func NewTaskService(r domain.TaskRepository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetTasks() []domain.Task {
	return s.repo.FindAll()
}

//Add creates a new task
func (s *Service) Add(task domain.Task) (domain.Task, error) {

	task.Id = uuid.New().String()
	//default due date is 3 days
	task.DueDate = time.Now().Add(72 * time.Hour).Format(time.RFC3339)

	err := s.repo.Add(task)
	if err != nil {
		return task, err
	}

	return task, nil

}

//FindByID returns a single task by ID
func (s *Service) FindByID(id string) (domain.Task, error) {

	var t domain.Task

	tasks := s.GetTasks()

	for _, t = range tasks {
		if t.Id == id {
			return t, nil
		}
	}

	return t, fmt.Errorf("the task with id %s doesn't exist", id)

}

func (s *Service) Update(t domain.Task) error {
	return s.repo.Update(t)
}

/* func (u *UserRepository) Update(user User) error {
	return u.updateUser(user, "name", "email")
}

func (u *UserRepository) Delete(id string) error {
	_, err := u.db.Exec(deleteUserSQL, id)
	return err
} */
