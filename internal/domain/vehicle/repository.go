package vehicle

import (
	"github.com/google/uuid"
	"go-ddd/internal/dtos"
)

type IVehicleRepository interface {
	GetAll() ([]dtos.VehicleOutput, error)
	Create(input dtos.VehicleCreateInput) error
	GetByID(uuid.UUID) (*Vehicle, error)
	Update(input *dtos.VehicleUpdateInput) error
	HardDelete(uuid.UUID) error
	SoftDelete(uuid.UUID) error
	UnDelete(uuid.UUID) error
	UnRelate(uuid.UUID) error
}
