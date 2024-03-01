package aggregate

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/moura95/go-ddd/internal/domain/vehicle"
	"time"
)

type DriverVehicleAggregate struct {
	Uuid          uuid.UUID
	Name          string
	Email         string
	TaxID         string
	DriverLicense string
	DateOfBirth   sql.NullString
	DeletedAt     sql.NullString
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Vehicles      []vehicle.Vehicle
}
