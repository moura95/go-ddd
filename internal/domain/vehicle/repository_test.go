package vehicle

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestGetAllVehicles_Successful(t *testing.T) {
	db, dbmock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	defer dbmock.ExpectClose()

	dbmock.MatchExpectationsInOrder(false)

	queryGetAll := "SELECT (.+) FROM vehicles"
	columns := []string{
		"uuid",
		"brand",
		"model",
		"year_of_manufacture",
		"license_plate",
		"color",
		"created_at",
		"update_at",
	}

	date := time.Now()

	rows := dbmock.NewRows(columns).
		AddRow("cc9fc54d-eca3-465c-acf1-0a6a373909ab", "Scania", "R500", 2020, "ABC123", "Blue", date, date).
		AddRow("16e308bb-4c07-4f32-bb36-7c92ced86df0", "Volvo", "FH16", 2019, "XYZ987", "Red", date, date)

	dbmock.ExpectQuery(queryGetAll).WithArgs().WillReturnRows(rows)
	repo := NewVehicleRepository(db, &zap.SugaredLogger{})

	data, err := repo.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(data))
}

func TestGetAllVehicles_Failed(t *testing.T) {
	db, dbmock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	defer dbmock.ExpectClose()

	dbmock.MatchExpectationsInOrder(false)

	queryGetAll := "SELECT (.+) FROM vehicles"

	mockError := errors.New("mocked error")

	dbmock.ExpectQuery(queryGetAll).WithArgs().WillReturnError(mockError)
	repo := NewVehicleRepository(db, &zap.SugaredLogger{})

	data, err := repo.GetAll()

	assert.Error(t, err)
	assert.Empty(t, data)
}

func TestCreateVehicle_SuccessfulWithAllRequiredFields(t *testing.T) {
	db, dbmock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	defer dbmock.ExpectClose()
	dbmock.MatchExpectationsInOrder(false)

	query := "INSERT INTO vehicles"

	vehicle := Vehicle{
		Brand:             "Volvo",
		Model:             "FH16",
		YearOfManufacture: 2019,
		LicensePlate:      "XYZ987",
		Color:             "Red",
	}

	dbmock.ExpectExec(query).
		WithArgs(
			vehicle.Brand,
			vehicle.Model,
			vehicle.YearOfManufacture,
			vehicle.LicensePlate,
			vehicle.Color).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewVehicleRepository(db, &zap.SugaredLogger{})

	err = repo.Create(vehicle)

	assert.NoError(t, err)
}

func TestGetByIDVehicle_Successful(t *testing.T) {
	db, dbmock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	defer dbmock.ExpectClose()
	dbmock.MatchExpectationsInOrder(false)

	queryGetByID := "SELECT (.+) FROM vehicles"

	columns := []string{
		"uuid",
		"brand",
		"model",
		"year_of_manufacture",
		"license_plate",
		"color",
		"created_at",
		"update_at",
	}

	UUID := uuid.MustParse("cc9fc54d-eca3-465c-acf1-0a6a373909ab")
	now := time.Now()

	expectedVehicle := &Vehicle{
		Uuid:              UUID,
		Brand:             "Scania",
		Model:             "R500",
		YearOfManufacture: 2020,
		LicensePlate:      "ABC123",
		Color:             "Blue",
		CreatedAt:         now,
		UpdatedAt:         now,
	}

	rows := dbmock.NewRows(columns).
		AddRow(UUID, "Scania", "R500", 2020, "ABC123", "Blue", now, now)

	dbmock.ExpectBegin()
	dbmock.ExpectQuery(queryGetByID).WithArgs().WillReturnRows(rows)
	dbmock.ExpectCommit()

	repo := NewVehicleRepository(db, &zap.SugaredLogger{})

	vehicle, err := repo.GetByID(UUID)
	assert.NoError(t, err)
	assert.Equal(t, expectedVehicle, vehicle)
}

func TestDeleteByIDVehicle_Successful(t *testing.T) {
	db, dbmock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	defer dbmock.ExpectClose()
	dbmock.MatchExpectationsInOrder(false)

	query := "DELETE FROM vehicles WHERE uuid = ?"

	UUID := uuid.MustParse("cc9fc54d-eca3-465c-acf1-0a6a373909ab")

	dbmock.ExpectExec(query).WithArgs(UUID).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewVehicleRepository(db, &zap.SugaredLogger{})

	err = repo.HardDelete(UUID)
	assert.NoError(t, err)
}
