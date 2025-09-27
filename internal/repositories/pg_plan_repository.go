package repositories

import (
	"sync"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PgPlanRepository struct {
	mu   sync.RWMutex
	DB   *gorm.DB
}

func NewPgPlanRepository(DB *gorm.DB) *PgPlanRepository {
	return &PgPlanRepository{
		DB: DB,
	}
}

func (r *PgPlanRepository) GetAll() ([]*Plan, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var plans []Plan
	result := r.DB.Find(&plans)
	if result.Error != nil {
		return nil, result.Error
	}

	var plansPointers []*Plan
		if len(plans) == 0 {
		return plansPointers, nil
	}

	for _, plan := range plans {
		plansPointers = append(plansPointers, &plan)
	}

	return plansPointers, nil
}

func (r *PgPlanRepository) GetById(id uuid.UUID) (*Plan, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var plan Plan
	result := r.DB.First(&plan, id)
		if result.Error != nil {
		return nil, result.Error
	}

	return &plan, nil
}

func (r *PgPlanRepository) Create(p Plan) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if result := r.DB.Create(&p); result.Error != nil {
        return result.Error
    }

	return nil
}
