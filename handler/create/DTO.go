package create

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type DTO struct {
	Method  string     `json:"method"`
	URL     string     `json:"url"`
	Headers HeadersDTO `json:"headers"`
}

type HeadersDTO struct {
	Authentication string `json:"authentication"`
}

func (dto *DTO) Validate() error {
	err := validation.ValidateStruct(dto,
		validation.Field(
			&dto.Method,
			validation.Required,
			validation.In("GET", "POST", "PUT", "PATCH", "DELETE"),
		),
		validation.Field(
			&dto.URL,
			validation.Required,
		),
	)
	if err != nil {
		return err
	}

	err = validation.ValidateStruct(&dto.Headers,
		validation.Field(
			&dto.Headers.Authentication,
			validation.Required,
		),
	)

	if err != nil {
		return err
	}

	return nil
}
