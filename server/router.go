package server

import (
	"les8/server/controller"
)

func (s *server) initRoutes(c *controller.Controller) {
	s.e.POST("/car", c.CreateCar)
	s.e.PUT("/car/:carID", c.UpdateCar)
	s.e.DELETE("/car/:carID", c.DeleteCar)
	s.e.GET("/car/:carID", c.GetCar)
	s.e.GET("/car/list", c.ListCars)
}
