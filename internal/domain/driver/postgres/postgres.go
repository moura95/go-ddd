package postgres

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go-ddd/internal/domain/driver"
	"go-ddd/internal/domain/vehicle"
	"go-ddd/internal/dtos"
	"go.uber.org/zap"
	"time"
)

type driverRepository struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewDriverRepository(db *sqlx.DB, log *zap.SugaredLogger) driver.IDriverRepository {
	return &driverRepository{db: db, logger: log}
}

func (r *driverRepository) GetAll() ([]dtos.DriverOuput, error) {
	var drivers []dtos.DriverOuput
	query := "SELECT * FROM drivers WHERE deleted_at is null"
	if err := r.db.Select(&drivers, query); err != nil {
		return []dtos.DriverOuput{}, err
	}
	return drivers, nil
}

func (r *driverRepository) Create(driver dtos.DriverCreateInput) error {
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

func (r *driverRepository) Subscribe(driverVehicle dtos.DriverVehicleInput) error {
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

func (r *driverRepository) UnSubscribe(driverVehicle dtos.DriverVehicleInput) error {
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

func (r *driverRepository) GetByID(uid uuid.UUID) (*dtos.DriverVehicle, error) {
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

	driver := &dtos.DriverVehicle{
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

func (r *driverRepository) Update(uuid uuid.UUID, driver *dtos.DriverUpdateInput) error {
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

func (r *driverRepository) HardDelete(uuid uuid.UUID) error {
	query := "DELETE FROM drivers WHERE uuid = :UUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"UUID": uuid})
	return err
}

func (r *driverRepository) SoftDelete(uuid uuid.UUID) error {
	query := "UPDATE drivers SET deleted_at=now() WHERE uuid = :UUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"UUID": uuid})
	return err
}

func (r *driverRepository) UnDelete(uuid uuid.UUID) error {
	query := "UPDATE drivers SET deleted_at=null WHERE uuid = :UUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"UUID": uuid})
	return err
}

func (r *driverRepository) UnRelate(driverUUID uuid.UUID) error {
	query := "DELETE FROM drivers_vehicles WHERE driver_uuid = :DriverUUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"DriverUUID": driverUUID})
	return err
}
