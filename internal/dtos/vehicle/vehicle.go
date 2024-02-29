package vehicle

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Output struct {
	Uuid              uuid.UUID      `db:"uuid"`
	Brand             string         `db:"brand"`
	Model             string         `db:"model"`
	YearOfManufacture uint           `db:"year_of_manufacture"`
	LicensePlate      string         `db:"license_plate"`
	Color             string         `db:"color"`
	DeletedAt         sql.NullString `db:"deleted_at"`
	CreatedAt         time.Time      `db:"created_at"`
	UpdatedAt         time.Time      `db:"update_at"`
}

type CreateInput struct {
	Brand             string `db:"brand"`
	Model             string `db:"model"`
	YearOfManufacture uint   `db:"year_of_manufacture"`
	LicensePlate      string `db:"license_plate"`
	Color             string `db:"color"`
}

type UpdateInput struct {
	Uuid              uuid.UUID `db:"uuid"`
	Brand             string    `db:"brand"`
	Model             string    `db:"model"`
	YearOfManufacture uint      `db:"year_of_manufacture"`
	LicensePlate      string    `db:"license_plate"`
	Color             string    `db:"color"`
}
