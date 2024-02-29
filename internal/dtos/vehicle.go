package dtos

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type VehicleOutput struct {
	Uuid              uuid.UUID      `db:"uuid"`
	Brand             string         `db:"brand"`
	Model             string         `db:"model"`
	YearOfManufacture uint           `db:"year_of_manufacture"`
	LicensePlate      string         `db:"license_plate"`
	Color             string         `db:"color"`
	DeletedAt         sql.NullString `db:"deleted_at"`
	CreatedAt         time.Time      `db:"created_at"`
	UpdatedAt         time.Time      `db:"updated_at"`
}

type VehicleCreateInput struct {
	Brand             string `db:"brand"`
	Model             string `db:"model"`
	YearOfManufacture uint   `db:"year_of_manufacture"`
	LicensePlate      string `db:"license_plate"`
	Color             string `db:"color"`
}

type VehicleUpdateInput struct {
	Uuid              uuid.UUID `db:"uuid"`
	Brand             string    `db:"brand"`
	Model             string    `db:"model"`
	YearOfManufacture uint      `db:"year_of_manufacture"`
	LicensePlate      string    `db:"license_plate"`
	Color             string    `db:"color"`
}
