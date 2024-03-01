package driver_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/moura95/go-ddd/internal/domain/driver/memory"
	dto "github.com/moura95/go-ddd/internal/dtos/driver"
	"github.com/stretchr/testify/assert"
)

type DriverServiceTest struct {
	repository memory.IDriverRepositoryMemory
}

func NewDriverServiceTest(repo memory.IDriverRepositoryMemory) *DriverServiceTest {
	return &DriverServiceTest{
		repository: repo,
	}
}

func TestCreateDriver(t *testing.T) {
	mockRepo := memory.NewDriverRepositoryMemory()
	service := NewDriverServiceTest(mockRepo)

	d := dto.CreateInput{
		Name:          "Driver Test",
		Email:         "test@example.com",
		TaxID:         "1234567890",
		DriverLicense: "XYZ12345",
	}

	err := service.repository.Create(d)
	if err != nil {
		t.Error("Failed to created")
	}
	assert.NoError(t, err)

}

func TestGetAll(t *testing.T) {
	mockRepo := memory.NewDriverRepositoryMemory()
	service := NewDriverServiceTest(mockRepo)

	drivers, err := service.repository.GetAll()
	if err != nil {
		t.Error("Failed to created")
	}
	assert.NoError(t, err)
	assert.Equal(t, drivers[0].Uuid, uuid.MustParse("61a218e4-7908-45d7-88bf-6226b53ab321"))
	assert.Equal(t, drivers[0].Name, "Driver 1")
	assert.Equal(t, drivers[0].Email, "driver1@example.com")
	assert.Equal(t, drivers[0].TaxID, "1234567890")
	assert.Equal(t, drivers[0].DriverLicense, "ABC12345")

	// assert record 2
	assert.Equal(t, drivers[1].Uuid, uuid.MustParse("ef9da75e-949f-4780-92b5-eda71618fc6c"))
	assert.Equal(t, drivers[1].Name, "Driver 2")
	assert.Equal(t, drivers[1].Email, "driver2@example.com")
	assert.Equal(t, drivers[1].TaxID, "9876543210")
	assert.Equal(t, drivers[1].DriverLicense, "XYZ98765")
}

func TestGetID(t *testing.T) {
	mockRepo := memory.NewDriverRepositoryMemory()
	service := NewDriverServiceTest(mockRepo)

	driver, err := service.repository.GetByID(uuid.MustParse("61a218e4-7908-45d7-88bf-6226b53ab321"))
	if err != nil {
		t.Error("Failed to get")
	}
	assert.NoError(t, err)
	assert.Equal(t, driver.Uuid, uuid.MustParse("61a218e4-7908-45d7-88bf-6226b53ab321"))
	assert.Equal(t, driver.Name, "Driver 1")
	assert.Equal(t, driver.Email, "driver1@example.com")
	assert.Equal(t, driver.TaxID, "1234567890")
	assert.Equal(t, driver.DriverLicense, "ABC12345")

}

func TestUpdate(t *testing.T) {
	mockRepo := memory.NewDriverRepositoryMemory()
	service := NewDriverServiceTest(mockRepo)

	uid := uuid.MustParse("61a218e4-7908-45d7-88bf-6226b53ab321")

	d := &dto.UpdateInput{
		Name:          "Drive Updated",
		Email:         "driver12345@example.com",
		TaxID:         "123456",
		DriverLicense: "ABC123451",
	}
	err := service.repository.Update(uid, d)
	if err != nil {
		t.Error("Failed to update")
	}
	assert.NoError(t, err)
}

func TestHardDelete(t *testing.T) {
	mockRepo := memory.NewDriverRepositoryMemory()
	service := NewDriverServiceTest(mockRepo)

	uid := uuid.MustParse("61a218e4-7908-45d7-88bf-6226b53ab321")
	err := service.repository.HardDelete(uid)
	if err != nil {
		t.Error("Failed to delete")
	}
	assert.NoError(t, err)
}

func TestSoftDelete(t *testing.T) {
	mockRepo := memory.NewDriverRepositoryMemory()
	service := NewDriverServiceTest(mockRepo)

	uid := uuid.MustParse("61a218e4-7908-45d7-88bf-6226b53ab321")
	err := service.repository.SoftDelete(uid)
	if err != nil {
		t.Error("Failed to delete")
	}
	assert.NoError(t, err)
}
