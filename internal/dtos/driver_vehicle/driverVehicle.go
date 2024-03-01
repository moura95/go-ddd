package driver_vehicle

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/moura95/go-ddd/internal/domain/vehicle"
)

type Input struct {
	DriverUUID  uuid.UUID `db:"driver_uuid"`
	VehicleUUID uuid.UUID `db:"vehicle_uuid"`
}

type Output struct {
	Uuid          uuid.UUID      `db:"uuid"`
	Name          string         `db:"name"`
	Email         string         `db:"email"`
	TaxID         string         `db:"tax_id"`
	DriverLicense string         `db:"driver_license"`
	DateOfBirth   sql.NullString `db:"date_of_birth"`
	DeletedAt     sql.NullString `db:"deleted_at"`
	CreatedAt     time.Time      `db:"created_at"`
	UpdatedAt     time.Time      `db:"update_at"`
	Vehicles      []vehicle.Vehicle
}
