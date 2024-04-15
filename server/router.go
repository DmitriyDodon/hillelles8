package server

import (
	"les8/server/controller"
)

func (s *server) initRoutes(c *controller.Controller) {
	s.e.POST("/car", c.CreateCar)
	s.e.PUT("/car/:carId", c.UpdateCar)
	s.e.DELETE("/car/:carId", c.DeleteCar)
	s.e.GET("/car/:carId", c.GetCar)
	s.e.GET("/car/list", c.ListCars)
}
