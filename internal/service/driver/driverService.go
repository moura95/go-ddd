package driver

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go-ddd/internal/aggregate"
	"go-ddd/internal/domain/driver"
	driver_dto "go-ddd/internal/dtos/driver"
	"go-ddd/internal/dtos/driver_vehicle"
	"go-ddd/internal/infra/cfg"
	"go.uber.org/zap"
)

type IDriverService interface {
	Create(driver driver_dto.CreateInput) error
	Subscribe(driverVehicle driver_vehicle.Input) error
	UnSubscribe(driverVehicle driver_vehicle.Input) error
	List() ([]driver_dto.Output, error)
	GetByID(uid uuid.UUID) (*aggregate.DriverVehicleAggregate, error)
	Update(driver driver_dto.UpdateInput) error
	SoftDelete(uid uuid.UUID) error
	UnDelete(uid uuid.UUID) error
	HardDelete(uid uuid.UUID) error
}

type driverService struct {
	database   *sqlx.DB
	repository driver.IDriverRepository
	config     cfg.Config
	logger     *zap.SugaredLogger
}

func NewDriverService(db *sqlx.DB, repo driver.IDriverRepository, cfg cfg.Config, log *zap.SugaredLogger) *driverService {
	return &driverService{
		database:   db,
		repository: repo,
		config:     cfg,
		logger:     log,
	}
}

func (d *driverService) Create(dto driver_dto.CreateInput) error {
	dr := driver_dto.CreateInput{
		Name:          dto.Name,
		Email:         dto.Email,
		TaxID:         dto.TaxID,
		DriverLicense: dto.DriverLicense,
		DateOfBirth:   dto.DateOfBirth,
	}
	err := d.repository.Create(dr)
	if err != nil {
		return fmt.Errorf("failed to create %s", err.Error())
	}
	return nil
}
func (d *driverService) Subscribe(driverVehicle driver_vehicle.Input) error {
	err := d.repository.Subscribe(driverVehicle)
	if err != nil {
		return fmt.Errorf("failed to create %s", err.Error())
	}
	return nil
}
func (d *driverService) UnSubscribe(driverVehicle driver_vehicle.Input) error {
	// unRelate driver before delete
	err := d.repository.UnSubscribe(driverVehicle)
	if err != nil {
		return fmt.Errorf("failed to hard delete driver %s", err.Error())
	}
	return nil
}

func (d *driverService) GetByID(uid uuid.UUID) (*aggregate.DriverVehicleAggregate, error) {
	driverOutput, err := d.repository.GetByID(uid)

	if err != nil {
		return &aggregate.DriverVehicleAggregate{}, fmt.Errorf("failed to get driver %s", err.Error())
	}
	return (*aggregate.DriverVehicleAggregate)(driverOutput), nil
}

func (d *driverService) List() ([]driver_dto.Output, error) {
	drivers, err := d.repository.GetAll()
	if err != nil {
		return []driver_dto.Output{}, fmt.Errorf("failed to list drivers %s", err.Error())
	}
	return drivers, nil
}

func (d *driverService) Update(dto driver_dto.UpdateInput) error {
	dr := driver_dto.UpdateInput{
		Uuid:          dto.Uuid,
		Name:          dto.Name,
		Email:         dto.Email,
		TaxID:         dto.TaxID,
		DriverLicense: dto.DriverLicense,
		DateOfBirth:   dto.DateOfBirth,
	}
	err := d.repository.Update(dr.Uuid, &dr)
	if err != nil {
		return fmt.Errorf("failed to update driver %s", err.Error())
	}
	return nil
}

func (d *driverService) SoftDelete(uid uuid.UUID) error {
	err := d.repository.SoftDelete(uid)
	if err != nil {
		return fmt.Errorf("failed to delete driver %s", err.Error())
	}
	return nil
}

func (d *driverService) UnDelete(uid uuid.UUID) error {
	err := d.repository.UnDelete(uid)
	if err != nil {
		return fmt.Errorf("failed to recover driver %s", err.Error())
	}
	return nil
}

func (d *driverService) HardDelete(uid uuid.UUID) error {
	// unRelate driver before delete
	err := d.repository.UnRelate(uid)
	if err != nil {
		return fmt.Errorf("failed to hard delete driver %s", err.Error())
	}

	err = d.repository.HardDelete(uid)
	if err != nil {
		return fmt.Errorf("failed to hardelete driver %s", err.Error())
	}
	return nil
}
