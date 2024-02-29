package dtos

import (
	"database/sql"
	"github.com/google/uuid"
	"go-ddd/internal/domain/vehicle"
	"time"
)

type DriverVehicleInput struct {
	DriverUUID  uuid.UUID `db:"driver_uuid"`
	VehicleUUID uuid.UUID `db:"vehicle_uuid"`
}

type DriverVehicle struct {
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
