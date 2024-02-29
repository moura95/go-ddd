package driver

import (
	"github.com/google/uuid"
	"go-ddd/internal/dtos"
)

type IDriverRepository interface {
	GetAll() ([]dtos.DriverOuput, error)
	Create(input dtos.DriverCreateInput) error
	Subscribe(driver dtos.DriverVehicleInput) error
	UnSubscribe(vehicle dtos.DriverVehicleInput) error
	GetByID(uuid.UUID) (*dtos.DriverVehicle, error)
	Update(uuid.UUID, *dtos.DriverUpdateInput) error
	HardDelete(uuid.UUID) error
	SoftDelete(uuid.UUID) error
	UnDelete(uuid.UUID) error
	UnRelate(uuid2 uuid.UUID) error
}
