package memory

import (
	"github.com/google/uuid"
	"go-ddd/internal/aggregate/driver_vehicle"
	"go-ddd/internal/domain/vehicle"
	"time"
)

type IVehicleRepositoryMock interface {
	GetAll() ([]vehicle.Vehicle, error)
	Create(vehicle.Vehicle) error
	Subscribe(aggregate.DriverVehicleAggregate) error
	UnSubscribe(aggregate.DriverVehicleAggregate) error
	GetByID(uuid.UUID) (*vehicle.Vehicle, error)
	Update(*vehicle.Vehicle) error
	HardDelete(uuid.UUID) error
	SoftDelete(uuid.UUID) error
	UnDelete(uuid.UUID) error
	UnRelate(vehicleUUID uuid.UUID) error
}

type VehicleRepositoryMock struct {
	vehicles []vehicle.Vehicle
}

func (v VehicleRepositoryMock) Subscribe(aggregate.DriverVehicleAggregate) error {
	//TODO implement me
	panic("implement me")
}

func (v VehicleRepositoryMock) UnSubscribe(aggregate.DriverVehicleAggregate) error {
	//TODO implement me
	panic("implement me")
}

func NewVehicleRepositoryMock() IVehicleRepositoryMock {
	return &VehicleRepositoryMock{vehicles: []vehicle.Vehicle{
		{
			Uuid:              uuid.MustParse("43ee3d4c-de06-4021-ab6f-ba8113418df9"),
			Brand:             "Scania",
			Model:             "R500",
			YearOfManufacture: uint(2020),
			LicensePlate:      "ABC123",
			Color:             "Blue",
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		},
		{
			Uuid:              uuid.MustParse("457a8df2-782f-4f22-8233-623b694096a1"),
			Brand:             "Volvo",
			Model:             "FH16",
			YearOfManufacture: uint(2019),
			LicensePlate:      "XYZ987",
			Color:             "Black",
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}}}
}

func (v VehicleRepositoryMock) GetAll() ([]vehicle.Vehicle, error) {
	return v.vehicles, nil
}

func (v VehicleRepositoryMock) Create(vehicle vehicle.Vehicle) error {
	v.vehicles = append(v.vehicles, vehicle)
	return nil
}

func (v VehicleRepositoryMock) GetByID(u uuid.UUID) (*vehicle.Vehicle, error) {
	for _, vehicle := range v.vehicles {
		if vehicle.Uuid == u {
			return &vehicle, nil
		}
	}
	return nil, nil
}

func (v VehicleRepositoryMock) Update(vehicle *vehicle.Vehicle) error {
	for i, ve := range v.vehicles {
		if ve.Uuid == vehicle.Uuid {
			v.vehicles[i] = *vehicle
			return nil
		}
	}
	return nil
}

func (v VehicleRepositoryMock) HardDelete(u uuid.UUID) error {
	for i, vehicle := range v.vehicles {
		if vehicle.Uuid == u {
			v.vehicles = append(v.vehicles[:i], v.vehicles[i+1:]...)
			return nil
		}
	}
	return nil
}

func (v VehicleRepositoryMock) SoftDelete(u uuid.UUID) error {
	for i, vehicle := range v.vehicles {
		if vehicle.Uuid == u {
			vehicle.DeletedAt.String = time.Now().String()
			v.vehicles[i] = vehicle
			return nil
		}
	}
	return nil
}

func (v VehicleRepositoryMock) UnDelete(u uuid.UUID) error {
	for i, vehicle := range v.vehicles {
		if vehicle.Uuid == u {
			vehicle.DeletedAt.String = ""
			v.vehicles[i] = vehicle
			return nil
		}
	}
	return nil
}

func (v VehicleRepositoryMock) UnRelate(vehicleUUID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
