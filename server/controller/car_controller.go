package controller

import (
	"fmt"
	"les8/server/httpmodels"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (con Controller) CreateCar(c echo.Context) error {
	req := new(httpmodels.CarCreateRequest)

	err := c.Bind(req)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, httpmodels.UnprocessableEntity)
	}

	err = req.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	query := fmt.Sprintf("insert into cars (id, color, price_in_cents, max_speed_mph, max_speed_kmp, vendor_name, model_name) values ('%s', '%s', %d, %d, %d, '%s', '%s')",
		uuid.Must(uuid.NewRandom()).String(),
		req.Color,
		req.PriceInCents,
		req.MaxSpeedMPH,
		req.MaxSpeedKMP,
		req.VendorName,
		req.ModelName,
	)

	_, err = con.dbConnection.Execute(query)

	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, httpmodels.ServerError)
	}

	return c.NoContent(http.StatusCreated)
}

func (con Controller) UpdateCar(c echo.Context) error {
	carId := c.Param("carId")

	req := new(httpmodels.CarCreateRequest)

	err := c.Bind(req)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, httpmodels.UnprocessableEntity)
	}

	err = req.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	query := fmt.Sprintf("update cars set color = '%s', price_in_cents = '%d', max_speed_mph = '%d', max_speed_kmp = '%d', vendor_name = '%s', model_name = '%s' where id = '%s'",
		req.Color,
		req.PriceInCents,
		req.MaxSpeedMPH,
		req.MaxSpeedKMP,
		req.VendorName,
		req.ModelName,
		carId,
	)

	_, err = con.dbConnection.Execute(query)

	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, httpmodels.ServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func (con Controller) DeleteCar(c echo.Context) error {
	carId := c.Param("carId")

	_, err := con.dbConnection.Execute(fmt.Sprintf("delete from cars where id = '%s'", carId))

	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, httpmodels.ServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func (con Controller) GetCar(c echo.Context) error {
	carId := c.Param("carId")
	carRow := con.dbConnection.QueryRow(fmt.Sprintf("select * from cars where id = '%s'", carId))

	car := new(httpmodels.CarResponse)

	err := carRow.Scan(&car.Id, &car.Color, &car.PriceInCents, &car.MaxSpeedMPH, &car.MaxSpeedKMP, &car.VendorName, &car.ModelName, &car.DateCreatedAt)

	if err != nil {
		return c.JSON(http.StatusNotFound, httpmodels.NotFoundError)
	}

	return c.JSON(http.StatusOK, car)
}

func (con Controller) ListCars(c echo.Context) error {
	carRows, err := con.dbConnection.Query(fmt.Sprintf("select * from cars"))

	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, httpmodels.ServerError)
	}

	cars := make([]httpmodels.CarResponse, 0)

	for carRows.Next() {
		car := new(httpmodels.CarResponse)

		err := carRows.Scan(&car.Id, &car.Color, &car.PriceInCents, &car.MaxSpeedMPH, &car.MaxSpeedKMP, &car.VendorName, &car.ModelName, &car.DateCreatedAt)

		if err != nil {
			// если произошла ошибка при парсинге строки из бд, тогда пропускаем такую entity
			continue
		}

		cars = append(cars, *car)
	}

	return c.JSON(http.StatusOK, cars)
}
