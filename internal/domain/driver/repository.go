package driver

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go-ddd/internal/aggregate/driver_vehicle"
	"go-ddd/internal/domain/vehicle"
	"go-ddd/internal/dtos"
	"go.uber.org/zap"
	"time"
)

type IDriverRepository interface {
	GetAll() ([]dtos.Driver, error)
	Create(Driver) error
	Subscribe(driver dtos.DriverVehicleInput) error
	UnSubscribe(vehicle dtos.DriverVehicleInput) error
	GetByID(uuid.UUID) (*aggregate.DriverVehicleAggregate, error)
	Update(uuid.UUID, *Driver) error
	HardDelete(uuid.UUID) error
	SoftDelete(uuid.UUID) error
	UnDelete(uuid.UUID) error
	UnRelate(uuid2 uuid.UUID) error
}

type DriverRepository struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewDriverRepository(db *sqlx.DB, log *zap.SugaredLogger) IDriverRepository {
	return &DriverRepository{db: db, logger: log}
}

func (r *DriverRepository) GetAll() ([]Driver, error) {
	var drivers []Driver
	query := "SELECT * FROM drivers WHERE deleted_at is null"
	if err := r.db.Select(&drivers, query); err != nil {
		return []Driver{}, err
	}
	return drivers, nil
}

func (r *DriverRepository) Create(driver Driver) error {
	query := `
        INSERT INTO drivers (name, email, tax_id, driver_license, date_of_birth)
        VALUES ($1, $2, $3, $4, $5)
    `
	args := []interface{}{
		driver.Name,
		driver.Email,
		driver.TaxID,
		driver.DriverLicense,
		driver.DateOfBirth,
	}
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *DriverRepository) Subscribe(driverVehicle dtos.DriverVehicleInput) error {
	query := `
        INSERT INTO drivers_vehicles (driver_uuid, vehicle_uuid)
        VALUES ($1, $2)
    `
	args := []interface{}{
		driverVehicle.VehicleUUID,
		driverVehicle.DriverUUID,
	}
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *DriverRepository) UnSubscribe(driverVehicle dtos.DriverVehicleInput) error {
	query := "DELETE FROM drivers_vehicles WHERE driver_uuid =$1 AND vehicle_uuid =$2"
	args := []interface{}{
		driverVehicle.VehicleUUID,
		driverVehicle.DriverUUID,
	}
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return err
}

func (r *DriverRepository) GetByID(uid uuid.UUID) (*aggregate.DriverVehicleAggregate, error) {
	var result []struct {
		DriverUUID    uuid.UUID      `database:"uuid"`
		DriverName    string         `database:"name"`
		DriverEmail   string         `database:"email"`
		DriverTaxID   string         `database:"tax_id"`
		DriverLicense string         `database:"driver_license"`
		DriverDOB     sql.NullString `database:"date_of_birth"`
		VehicleUUID   uuid.UUID      `database:"uuid"`
		VehicleBrand  string         `database:"brand"`
		VehicleModel  string         `database:"model"`
		VehicleYear   uint           `database:"year_of_manufacture"`
		VehicleColor  string         `database:"color"`
	}

	query := `
		SELECT d.uuid, d.name, d.email, d.tax_id, d.driver_license, d.date_of_birth,
		       v.uuid, v.brand , v.model,
		       v.year_of_manufacture , v.color
		FROM drivers AS d
		LEFT JOIN drivers_vehicles AS dv ON d.uuid = dv.driver_uuid
		LEFT JOIN vehicles AS v ON v.uuid = dv.vehicle_uuid
		WHERE d.uuid = $1
	`

	err := r.db.Select(&result, query, uid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	driver := &aggregate.DriverVehicleAggregate{
		Uuid:          result[0].DriverUUID,
		Name:          result[0].DriverName,
		Email:         result[0].DriverEmail,
		TaxID:         result[0].DriverTaxID,
		DriverLicense: result[0].DriverLicense,
		DateOfBirth:   result[0].DriverDOB,
		Vehicles:      make([]vehicle.Vehicle, 0, len(result)),
	}

	for _, res := range result {
		vehicle := vehicle.Vehicle{
			Uuid:              res.VehicleUUID,
			Brand:             res.VehicleBrand,
			Model:             res.VehicleModel,
			YearOfManufacture: res.VehicleYear,
			Color:             res.VehicleColor,
		}
		driver.Vehicles = append(driver.Vehicles, vehicle)
	}

	return driver, nil
}

func (r *DriverRepository) Update(uuid uuid.UUID, driver *Driver) error {
	query := `
        UPDATE drivers 
        SET name=$2, tax_id=$3, driver_license=$4, date_of_birth=$5, update_at=$6
    	WHERE uuid= $1`

	args := []interface{}{
		uuid,
		driver.Name,
		driver.TaxID,
		driver.DriverLicense,
		driver.DateOfBirth,
		time.Now(),
	}
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *DriverRepository) HardDelete(uuid uuid.UUID) error {
	query := "DELETE FROM drivers WHERE uuid = :UUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"UUID": uuid})
	return err
}

func (r *DriverRepository) SoftDelete(uuid uuid.UUID) error {
	query := "UPDATE drivers SET deleted_at=now() WHERE uuid = :UUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"UUID": uuid})
	return err
}

func (r *DriverRepository) UnDelete(uuid uuid.UUID) error {
	query := "UPDATE drivers SET deleted_at=null WHERE uuid = :UUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"UUID": uuid})
	return err
}

func (r *DriverRepository) UnRelate(driverUUID uuid.UUID) error {
	query := "DELETE FROM drivers_vehicles WHERE driver_uuid = :DriverUUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"DriverUUID": driverUUID})
	return err
}
