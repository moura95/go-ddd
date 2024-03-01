package driver

import (
	"github.com/google/uuid"
	"github.com/moura95/go-ddd/internal/dtos/driver"
	"github.com/moura95/go-ddd/internal/dtos/driver_vehicle"
)

type IDriverRepository interface {
	GetAll() ([]driver.Output, error)
	Create(input driver.CreateInput) error
	Subscribe(driver driver_vehicle.Input) error
	UnSubscribe(vehicle driver_vehicle.Input) error
	GetByID(uuid.UUID) (*driver_vehicle.Output, error)
	Update(uuid.UUID, *driver.UpdateInput) error
	HardDelete(uuid.UUID) error
	SoftDelete(uuid.UUID) error
	UnDelete(uuid.UUID) error
	UnRelate(uuid2 uuid.UUID) error
}
