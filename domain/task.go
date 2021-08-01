package domain

type Task struct {
	Id      string `json:"id,omitempty"`
	Topic   string `json:"topic" validate:"required,min=3"`
	Status  bool   `json:"status" validate:"required"`
	DueDate string `json:"date,omitempty"`
}

/* func (t *Task) Validate() error {
	return validator.New().Struct(t)
} */

// Interface for Task repo
type TaskRepository interface {
	Add(u Task) error
	FindByID(id string) (Task, error)
	FindAll() []Task
	Update(t Task) error
}
