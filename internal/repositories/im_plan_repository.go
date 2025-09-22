package repositories

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type InMemoryPlanRepository struct {
	mu   sync.RWMutex
	data []Plan
}

func NewInMemoryPlanRepository() *InMemoryPlanRepository {
	return &InMemoryPlanRepository{
		data: make([]Plan, 0),
	}
}

func (r *InMemoryPlanRepository) GetAll() ([]*Plan, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]*Plan, 0, len(r.data))
	for _, plan := range r.data {
		out = append(out, &plan)
	}

	return out, nil
}

func (r *InMemoryPlanRepository) GetById(id uuid.UUID) (*Plan, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, plan := range r.data {
		if plan.Id == id {
			return &plan, nil
		}
	}

	return nil, errors.New("plan not found")
}

func (r *InMemoryPlanRepository) Create(p Plan) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.data = append(r.data, p)

	return nil
}
