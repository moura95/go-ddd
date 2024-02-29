package memory

import (
	"database/sql"
	"github.com/google/uuid"
	"go-ddd/internal/aggregate"
	"go-ddd/internal/domain/driver"
	dto "go-ddd/internal/dtos/driver"
	"go-ddd/internal/dtos/driver_vehicle"

	"time"
)

func (m *driverRepositoryMemory) GetAll() ([]driver.Driver, error) {

	return m.drivers, nil
}

func (m *driverRepositoryMemory) Create(dto dto.CreateInput) error {
	d := driver.NewDriver(dto.Name, dto.Email, dto.TaxID, dto.DriverLicense, dto.DateOfBirth.String)
	m.drivers = append(m.drivers, *d)
	return nil
}

func (m *driverRepositoryMemory) GetByID(u uuid.UUID) (*aggregate.DriverVehicleAggregate, error) {
	for _, d := range m.drivers {
		if d.Uuid == u {

			return &aggregate.DriverVehicleAggregate{
				Uuid:          d.Uuid,
				Name:          d.Name,
				Email:         d.Email,
				TaxID:         d.TaxID,
				DriverLicense: d.DriverLicense,
				DateOfBirth:   d.DateOfBirth,
				DeletedAt:     sql.NullString{},
				CreatedAt:     time.Time{},
				UpdatedAt:     time.Time{},
				Vehicles:      nil,
			}, nil
		}
	}
	return nil, nil
}

func (m *driverRepositoryMemory) Update(u uuid.UUID, dto *dto.UpdateInput) error {
	for i, d := range m.drivers {
		if d.Uuid == u {
			m.drivers[i] = *driver.NewDriver(dto.Name, dto.Email, dto.TaxID, dto.DriverLicense, dto.DateOfBirth.String)
			return nil
		}
	}
	return nil
}

func (m *driverRepositoryMemory) HardDelete(u uuid.UUID) error {
	for i, driver := range m.drivers {
		if driver.Uuid == u {
			m.drivers = append(m.drivers[:i], m.drivers[i+1:]...)
			return nil
		}
	}
	return nil
}

func (m *driverRepositoryMemory) SoftDelete(u uuid.UUID) error {
	for i, driver := range m.drivers {
		if driver.Uuid == u {
			driver.DeletedAt.String = time.Now().String()
			m.drivers[i] = driver
			return nil
		}
	}
	return nil
}

func (m *driverRepositoryMemory) UnDelete(u uuid.UUID) error {
	for i, driver := range m.drivers {
		if driver.Uuid == u {
			driver.DeletedAt.String = ""
			m.drivers[i] = driver
			return nil
		}
	}
	return nil
}

func (m *driverRepositoryMemory) Subscribe(driver driver_vehicle.Input) error {
	//TODO implement me
	panic("implement me")
}

func (m *driverRepositoryMemory) UnSubscribe(vehicle driver_vehicle.Input) error {
	//TODO implement me
	panic("implement me")
}

func (m *driverRepositoryMemory) UnRelate(uuid2 uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

type IDriverRepositoryMemory interface {
	GetAll() ([]driver.Driver, error)
	Create(input dto.CreateInput) error
	Subscribe(driver driver_vehicle.Input) error
	UnSubscribe(vehicle driver_vehicle.Input) error
	GetByID(uuid.UUID) (*aggregate.DriverVehicleAggregate, error)
	Update(uuid.UUID, *dto.UpdateInput) error
	HardDelete(uuid.UUID) error
	SoftDelete(uuid.UUID) error
	UnDelete(uuid.UUID) error
	UnRelate(uuid2 uuid.UUID) error
}

type driverRepositoryMemory struct {
	drivers []driver.Driver
}

func NewDriverRepositoryMemory() IDriverRepositoryMemory {
	return &driverRepositoryMemory{
		drivers: []driver.Driver{
			{
				Uuid:          uuid.MustParse("61a218e4-7908-45d7-88bf-6226b53ab321"),
				Name:          "Driver 1",
				Email:         "driver1@example.com",
				TaxID:         "1234567890",
				DriverLicense: "ABC12345",
				DateOfBirth:   sql.NullString{String: "1990-01-01", Valid: true},
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			},
			{
				Uuid:          uuid.MustParse("ef9da75e-949f-4780-92b5-eda71618fc6c"),
				Name:          "Driver 2",
				Email:         "driver2@example.com",
				TaxID:         "9876543210",
				DriverLicense: "XYZ98765",
				DateOfBirth:   sql.NullString{String: "1985-02-15", Valid: true},
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			},
		},
	}
}
