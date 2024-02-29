package driver

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

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
}

type CreateInput struct {
	Name          string
	Email         string
	TaxID         string
	DriverLicense string
	DateOfBirth   sql.NullString
}

type UpdateInput struct {
	Uuid          uuid.UUID
	Name          string
	Email         string
	TaxID         string
	DriverLicense string
	DateOfBirth   sql.NullString
}
