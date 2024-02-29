package vehicle

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go-ddd/internal/domain/vehicle"
	"go-ddd/internal/infra/cfg"
	"go.uber.org/zap"
)

type IVehicleService interface {
	Create(vehicle vehicle.Vehicle) error
	List() ([]vehicle.Vehicle, error)
	GetByID(uid uuid.UUID) (*vehicle.Vehicle, error)
	Update(vehicle vehicle.Vehicle) error
	SoftDelete(uid uuid.UUID) error
	UnDelete(uid uuid.UUID) error
	HardDelete(uid uuid.UUID) error
}

type vehicleService struct {
	database   *sqlx.DB
	repository vehicle.IVehicleRepository
	config     cfg.Config
	logger     *zap.SugaredLogger
}

func NewVehicleService(db *sqlx.DB, repo vehicle.IVehicleRepository, cfg cfg.Config, log *zap.SugaredLogger) *vehicleService {
	return &vehicleService{
		database:   db,
		repository: repo,
		config:     cfg,
		logger:     log,
	}
}

func (v *vehicleService) Create(vehicle vehicle.Vehicle) error {
	err := v.repository.Create(vehicle)
	if err != nil {
		return fmt.Errorf("failed to create %s", err.Error())
	}
	return nil
}

func (v *vehicleService) List() ([]vehicle.Vehicle, error) {
	vehicles, err := v.repository.GetAll()
	if err != nil {
		return []vehicle.Vehicle{}, fmt.Errorf("failed to get %s", err.Error())

	}
	return vehicles, nil
}

func (v *vehicleService) GetByID(uid uuid.UUID) (*vehicle.Vehicle, error) {
	ve, err := v.repository.GetByID(uid)
	if err != nil {
		return &vehicle.Vehicle{}, fmt.Errorf("failed to get %s", err.Error())

	}
	return ve, nil

}

func (v *vehicleService) Update(vehicle vehicle.Vehicle) error {
	err := v.repository.Update(&vehicle)
	if err != nil {
		return fmt.Errorf("failed to update %s", err.Error())
	}
	return nil
}

func (v *vehicleService) SoftDelete(uid uuid.UUID) error {
	err := v.repository.SoftDelete(uid)
	if err != nil {
		return fmt.Errorf("failed to delete")
	}
	return nil
}
func (v *vehicleService) UnDelete(uid uuid.UUID) error {
	err := v.repository.UnDelete(uid)
	if err != nil {
		return fmt.Errorf("failed to un delete %s", err.Error())
	}
	return nil
}

func (v *vehicleService) HardDelete(uid uuid.UUID) error {
	// unRelate driver before delete
	err := v.repository.UnRelate(uid)
	if err != nil {
		return fmt.Errorf("failed to delete %s", err.Error())
	}

	err = v.repository.HardDelete(uid)
	if err != nil {
		return fmt.Errorf("failed to  delete %s", err.Error())
	}
	return nil
}
