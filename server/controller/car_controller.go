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

	_, err = con.dbConnection.Execute("insert into cars (id, color, price_in_cents, max_speed_mph, max_speed_kmp, vendor_name, model_name) values (?, ?, ?, ?, ?, ?, ?)",
		uuid.Must(uuid.NewRandom()).String(),
		req.Color,
		req.PriceInCents,
		req.MaxSpeedMPH,
		req.MaxSpeedKMP,
		req.VendorName,
		req.ModelName,
	)

	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, httpmodels.ServerError)
	}

	return c.NoContent(http.StatusCreated)
}

func (con Controller) UpdateCar(c echo.Context) error {
	carID := c.Param("carID")

	req := new(httpmodels.CarCreateRequest)

	err := c.Bind(req)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, httpmodels.UnprocessableEntity)
	}

	err = req.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, err = con.dbConnection.Execute("update cars set color = ?, price_in_cents = ?, max_speed_mph = ?, max_speed_kmp = ?, vendor_name = ?, model_name = ? where id = ?",
		req.Color,
		req.PriceInCents,
		req.MaxSpeedMPH,
		req.MaxSpeedKMP,
		req.VendorName,
		req.ModelName,
		carID,
	)

	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, httpmodels.ServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func (con Controller) DeleteCar(c echo.Context) error {
	carID := c.Param("carID")

	_, err := con.dbConnection.Execute("delete from cars where id = ?", carID)

	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, httpmodels.ServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func (con Controller) GetCar(c echo.Context) error {
	carID := c.Param("carID")
	carRow := con.dbConnection.QueryRow("select * from cars where id = ?", carID)

	car := new(httpmodels.CarResponse)

	err := carRow.Scan(&car.Id, &car.Color, &car.PriceInCents, &car.MaxSpeedMPH, &car.MaxSpeedKMP, &car.VendorName, &car.ModelName, &car.DateCreatedAt)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, httpmodels.NotFoundError)
	}

	return c.JSON(http.StatusOK, car)
}

func (con Controller) ListCars(c echo.Context) error {
	carRows, err := con.dbConnection.Query("select * from cars")

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
