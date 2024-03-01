package vehicle

import (
	"github.com/google/uuid"
	"github.com/moura95/go-ddd/internal/dtos/vehicle"
)

type IVehicleRepository interface {
	GetAll() ([]vehicle.Output, error)
	Create(input vehicle.CreateInput) error
	GetByID(uuid.UUID) (*vehicle.Output, error)
	Update(input *vehicle.UpdateInput) error
	HardDelete(uuid.UUID) error
	SoftDelete(uuid.UUID) error
	UnDelete(uuid.UUID) error
	UnRelate(uuid.UUID) error
}
