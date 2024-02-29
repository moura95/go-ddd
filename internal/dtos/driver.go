package dtos

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Driver struct {
	Uuid          uuid.UUID      `database:"uuid"`
	Name          string         `database:"name"`
	Email         string         `database:"email"`
	TaxID         string         `database:"tax_id"`
	DriverLicense string         `database:"driver_license"`
	DateOfBirth   sql.NullString `database:"date_of_birth"`
	DeletedAt     sql.NullString `database:"deleted_at"`
	CreatedAt     time.Time      `database:"created_at"`
	UpdatedAt     time.Time      `database:"update_at"`
}

type DriverCreateInput struct {
	Name          string
	Email         string
	TaxID         string
	DriverLicense string
	DateOfBirth   sql.NullString
}

type DriverUpdateInput struct {
	Uuid          uuid.UUID
	Name          string
	Email         string
	TaxID         string
	DriverLicense string
	DateOfBirth   sql.NullString
}
