package gin

import (
	driver2 "go-ddd/internal/domain/driver"
	vehicle2 "go-ddd/internal/domain/vehicle"
	driver_router "go-ddd/internal/infra/http/gin/router/driver"
	vehicle_router "go-ddd/internal/infra/http/gin/router/vehicle"
	"go-ddd/internal/service/driver"
	"go-ddd/internal/service/vehicle"
	"go.uber.org/zap"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) createRoutesV1(router *gin.Engine, log *zap.SugaredLogger) {
	router.GET("/healthz", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	routes := router.Group("/")
	// Instance Driver Repository
	driverRepository := driver2.NewDriverRepository(s.store, log)
	// Instance Driver Service
	driverService := driver.NewDriverService(s.store, driverRepository, *s.config, log)

	// Instance VehicleRouter Repository
	vehicleRepository := vehicle2.NewVehicleRepository(s.store, log)
	// Instance VehicleRouter Service
	vehicleService := vehicle.NewVehicleService(s.store, vehicleRepository, *s.config, log)

	vehicle_router.NewVehicleRouter(vehicleService, log).SetupVehicleRoute(routes)
	driver_router.NewDriverRouter(driverService, log).SetupDriverRoute(routes)
}
