package driver

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	aggregate "go-ddd/internal/aggregate/driver_vehicle"
	"go-ddd/internal/domain/driver"
	"go-ddd/internal/dtos"
	"go-ddd/internal/infra/cfg"
	"go.uber.org/zap"
)

type IDriverService interface {
	Create(driver driver.Driver) error
	Subscribe(driverVehicle dtos.DriverVehicleInput) error
	UnSubscribe(driverVehicle dtos.DriverVehicleInput) error
	List() ([]driver.Driver, error)
	GetByID(uid uuid.UUID) (*aggregate.DriverVehicleAggregate, error)
	Update(driver dtos.DriverUpdateInput) error
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

func (d *driverService) Create(dto dtos.DriverCreateInput) error {
	dr := driver.Driver{
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
func (d *driverService) Subscribe(driverVehicle dtos.DriverVehicleInput) error {
	err := d.repository.Subscribe(driverVehicle)
	if err != nil {
		return fmt.Errorf("failed to create %s", err.Error())
	}
	return nil
}
func (d *driverService) UnSubscribe(driverVehicle dtos.DriverVehicleInput) error {
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

func (d *driverService) List() ([]dtos.DriverOuput, error) {
	drivers, err := d.repository.GetAll()
	if err != nil {
		return []dtos.DriverOuput{}, fmt.Errorf("failed to list drivers %s", err.Error())
	}
	return drivers, nil
}

func (d *driverService) Update(dto dtos.DriverUpdateInput) error {
	dr := driver.Driver{
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
