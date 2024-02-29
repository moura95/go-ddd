package vehicle

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"time"
)

type Vehicle struct {
	Uuid              uuid.UUID      `database:"uuid"`
	Brand             string         `database:"brand"`
	Model             string         `database:"model"`
	YearOfManufacture uint           `database:"year_of_manufacture"`
	LicensePlate      string         `database:"license_plate"`
	Color             string         `database:"color"`
	DeletedAt         sql.NullString `database:"deleted_at"`
	CreatedAt         time.Time      `database:"created_at"`
	UpdatedAt         time.Time      `database:"update_at"`
}

func NewVehicle(brand, model, licensePlate, color string, yearOfManufacture uint) *Vehicle {
	return &Vehicle{
		Uuid:              uuid.New(),
		Brand:             brand,
		Model:             model,
		YearOfManufacture: yearOfManufacture,
		LicensePlate:      licensePlate,
		Color:             color,
	}
}

func (v *Vehicle) validate() error {
	if v.Brand == "" {
		return errors.New("invalid brand")
	}
	if v.Model == "" {
		return errors.New("invalid model")
	}
	if v.LicensePlate == "" {
		return errors.New("invalid license plate")
	}
	return nil

}
