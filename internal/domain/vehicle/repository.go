package vehicle

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"time"
)

type IVehicleRepository interface {
	GetAll() ([]Vehicle, error)
	Create(Vehicle) error
	GetByID(uuid.UUID) (*Vehicle, error)
	Update(*Vehicle) error
	HardDelete(uuid.UUID) error
	SoftDelete(uuid.UUID) error
	UnDelete(uuid.UUID) error
	UnRelate(uuid.UUID) error
}

type vehicleRepository struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewVehicleRepository(db *sqlx.DB, log *zap.SugaredLogger) IVehicleRepository {
	return &vehicleRepository{db: db, logger: log}
}

func (r *vehicleRepository) GetAll() ([]Vehicle, error) {
	var vehicles []Vehicle
	query := "SELECT * FROM vehicles WHERE deleted_at is null"
	if err := r.db.Select(&vehicles, query); err != nil {
		return nil, err
	}
	return vehicles, nil
}

func (r *vehicleRepository) Create(vehicle Vehicle) error {
	query := `
        INSERT INTO vehicles (brand, model, year_of_manufacture, license_plate, color)
        VALUES ($1, $2, $3, $4, $5)
    `
	args := []interface{}{
		vehicle.Brand,
		vehicle.Model,
		vehicle.YearOfManufacture,
		vehicle.LicensePlate,
		vehicle.Color,
	}
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
func (r *vehicleRepository) GetByID(uuid uuid.UUID) (*Vehicle, error) {
	var vehicle Vehicle
	err := r.db.Get(&vehicle, "SELECT * FROM vehicles WHERE uuid = $1", uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Driver not found
		}
		return nil, err
	}
	return &vehicle, nil
}

func (r *vehicleRepository) Update(vehicle *Vehicle) error {
	query := `
        UPDATE vehicles 
        SET brand=$2, model=$3, year_of_manufacture=$4, license_plate=$5, color=$6, update_at=$7
        WHERE uuid=$1
    `
	args := []interface{}{
		vehicle.Uuid,
		vehicle.Brand,
		vehicle.Model,
		vehicle.YearOfManufacture,
		vehicle.LicensePlate,
		vehicle.Color,
		time.Now(),
	}
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *vehicleRepository) HardDelete(uuid uuid.UUID) error {
	query := "DELETE FROM vehicles WHERE uuid = :UUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"UUID": uuid})
	return err
}
func (r *vehicleRepository) SoftDelete(uuid uuid.UUID) error {
	query := "UPDATE vehicles SET deleted_at=now() WHERE uuid = :UUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"UUID": uuid})
	return err
}

func (r *vehicleRepository) UnDelete(uuid uuid.UUID) error {
	query := "UPDATE vehicles SET deleted_at=null WHERE uuid = :UUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"UUID": uuid})
	return err
}

func (r *vehicleRepository) UnRelate(vehicleUUID uuid.UUID) error {
	query := "DELETE FROM drivers_vehicles WHERE vehicle_uuid = :VehicleUUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"VehicleUUID": vehicleUUID})
	return err
}
