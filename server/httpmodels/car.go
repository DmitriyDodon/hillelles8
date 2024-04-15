package httpmodels

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type CarCreateRequest struct {
	Color        string `json:"color" validate:"required,iscolor,max=255"`
	PriceInCents int    `json:"price_in_cents" validate:"required,gte=0"`
	MaxSpeedMPH  int    `json:"max_speed_mph" validate:"gte=0"`
	MaxSpeedKMP  int    `json:"max_speed_kmp" validate:"required,gte=0"`
	VendorName   string `json:"vendor_name" validate:"required,min=2,max=255"`
	ModelName    string `json:"model_name" validate:"required,min=2,max=255"`
}

type CarResponse struct {
	Id            string    `json:"id"`
	DateCreatedAt time.Time `json:"date_created_at"`
	CarCreateRequest
}

func (ccr *CarCreateRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(ccr)
	if err != nil {
		return err

	}

	return nil
}
