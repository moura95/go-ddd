package dtos

import "github.com/google/uuid"

type DriverVehicleInput struct {
	DriverUUID  uuid.UUID `database:"driver_uuid"`
	VehicleUUID uuid.UUID `database:"vehicle_uuid"`
}
