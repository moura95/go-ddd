package driver

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestGetAll_Successful(t *testing.T) {
	db, dbmock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	defer dbmock.ExpectClose()

	dbmock.MatchExpectationsInOrder(false)

	queryGetAll := "SELECT (.+) FROM drivers"
	columns := []string{
		"uuid",
		"name",
		"email",
		"tax_id",
		"driver_license",
		"date_of_birth",
		"created_at",
		"update_at",
	}

	date := time.Now()

	rows := dbmock.NewRows(columns).
		AddRow("cc9fc54d-eca3-465c-acf1-0a6a373909ab", "Motorista 1", "motorista1@example.com", "12345678901", "ABC12345", "", date, date).
		AddRow("16e308bb-4c07-4f32-bb36-7c92ced86df0", "Motorista 2", "motorista2@example.com", "23456789012", "XYZ54321", "", date, date)

	dbmock.ExpectQuery(queryGetAll).WithArgs().WillReturnRows(rows)
	repo := NewDriverRepository(db, &zap.SugaredLogger{})

	data, err := repo.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(data))
}

func TestGetAll_Failed(t *testing.T) {
	db, dbmock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	defer dbmock.ExpectClose()

	dbmock.MatchExpectationsInOrder(false)

	queryGetAll := "SELECT (.+) FROM drivers"

	mockError := errors.New("mocked error")

	dbmock.ExpectQuery(queryGetAll).WithArgs().WillReturnError(mockError)
	repo := NewDriverRepository(db, &zap.SugaredLogger{})

	data, err := repo.GetAll()

	assert.Error(t, err)
	assert.Empty(t, data)
}

func TestCreate_SuccessfulWithAllRequiredFields(t *testing.T) {
	db, dbmock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	defer dbmock.ExpectClose()
	dbmock.MatchExpectationsInOrder(false)

	query := "INSERT INTO drivers"

	driver := Driver{
		Name:          "John Doe",
		Email:         "johndoe@example.com",
		TaxID:         "1234567890",
		DriverLicense: "ABC123",
		DateOfBirth:   sql.NullString{String: "1990-01-01"},
	}

	dbmock.ExpectExec(query).
		WithArgs(
			driver.Name,
			driver.Email,
			driver.TaxID,
			driver.DriverLicense,
			driver.DateOfBirth).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewDriverRepository(db, &zap.SugaredLogger{})

	err = repo.Create(driver)

	assert.NoError(t, err)
}

func TestUpdate_Successful(t *testing.T) {
	db, dbmock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	defer dbmock.ExpectClose()
	dbmock.MatchExpectationsInOrder(false)

	query := "UPDATE drivers"

	driver := Driver{
		Name:          "John Doe",
		Email:         "johndoe@example.com",
		TaxID:         "1234567890",
		DriverLicense: "ABC123",
		DateOfBirth:   sql.NullString{String: "1990-01-01", Valid: true},
	}

	dbmock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewDriverRepository(db, &zap.SugaredLogger{})
	uID := uuid.MustParse("cc9fc54d-eca3-465c-acf1-0a6a373909ab")
	err = repo.Update(uID, &driver)

	assert.NoError(t, err)
}

func TestDeleteByID_Successful(t *testing.T) {
	db, dbmock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	defer dbmock.ExpectClose()
	dbmock.MatchExpectationsInOrder(false)

	query := "DELETE FROM drivers WHERE uuid = ?"

	UUID := uuid.MustParse("cc9fc54d-eca3-465c-acf1-0a6a373909ab")

	dbmock.ExpectExec(query).WithArgs(UUID).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewDriverRepository(db, &zap.SugaredLogger{})

	err = repo.HardDelete(UUID)
	assert.NoError(t, err)
}
