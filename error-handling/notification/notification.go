package notification

import (
	"time"
)

type Error struct {
	Name      string      `json:"name"`
	Message   string      `json:"message"`
	Timestamp time.Time   `json:"timestamp"`
	Context   interface{} `json:"context"`
}

type Notification struct {
	Errors []Error `json:"errors"`
}

func NewNotification() *Notification {
	return &Notification{Errors: []Error{}}
}

func (n *Notification) AddError(name, message string, context interface{}) {
	n.Errors = append(n.Errors, Error{
		Name:      name,
		Message:   message,
		Timestamp: time.Now(),
		Context:   context,
	})
}

func (n *Notification) HasErrors() bool {
	return len(n.Errors) > 0
}

func (n *Notification) GetErrors() []Error {
	return n.Errors
}
