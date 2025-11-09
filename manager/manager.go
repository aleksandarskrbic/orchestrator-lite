package manager

import (
	"github.com/aleksandarskrbic/orchestrator-lite/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pending        queue.Queue
	TaskDb         map[string][]*task.Task
	EventDb        map[string][]*task.TaskEvent
	Workers        []string
	WorkersTaskMap map[string][]uuid.UUID
	TaskWorkerMap  map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {

}

func (m *Manager) UpdateTasks() {

}

func (m *Manager) SendWork() {

}
