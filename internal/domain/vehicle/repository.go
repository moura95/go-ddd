package vehicle

import (
	"github.com/google/uuid"
)

type IVehicleRepository interface {
	GetAll() ([]Vehicle, error)
	Create(input Vehicle) error
	GetByID(uuid.UUID) (*Vehicle, error)
	Update(input *Vehicle) error
	HardDelete(uuid.UUID) error
	SoftDelete(uuid.UUID) error
	UnDelete(uuid.UUID) error
	UnRelate(uuid.UUID) error
}
