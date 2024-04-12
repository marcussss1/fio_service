package validation

import (
	"github.com/google/uuid"
	"github.com/marcussss1/fio_service/internal/models"
	"github.com/marcussss1/fio_service/internal/pkg/e"
)

func IsUUID(str string) error {
	_, err := uuid.Parse(str)
	if err != nil {
		return e.ErrInvalidPeopleID
	}

	return nil
}

func ValidateGetPeopleRequest(req models.GetPeopleRequest) error {
	if err := validateName(req.Name); err != nil {
		return err
	}
	if err := validateSurname(req.Surname); err != nil {
		return err
	}
	if err := validatePatronymic(req.Patronymic); err != nil {
		return err
	}
	if err := validateAge(req.StartAge); err != nil {
		return err
	}
	if err := validateAge(req.EndAge); err != nil {
		return err
	}
	if err := validateGender(req.Gender); req.Gender != "" && err != nil {
		return err
	}
	if err := validateNationality(req.Nationality); err != nil {
		return err
	}

	return nil
}

func ValidateAbbreviatedPeopleRequest(req models.AbbreviatedPeopleRequest) error {
	if err := validateName(req.Name); err != nil {
		return err
	}
	if err := validateSurname(req.Surname); err != nil {
		return err
	}
	if err := validatePatronymic(req.Patronymic); err != nil {
		return err
	}

	return nil
}

func validateName(name string) error {
	if len(name) > 50 {
		return e.ErrLongName
	}

	return nil
}

func validateSurname(surname string) error {
	if len(surname) > 50 {
		return e.ErrLongSurname
	}

	return nil
}

func validatePatronymic(patronymic string) error {
	if len(patronymic) > 50 {
		return e.ErrLongPatronymic
	}

	return nil
}

func validateAge(age int) error {
	if age < 0 || age > 150 {
		return e.ErrInvalidAge
	}

	return nil
}

func validateGender(gender string) error {
	if gender == "male" || gender == "female" {
		return nil
	}

	return e.ErrInvalidGender
}

func validateNationality(nationality string) error {
	if len(nationality) > 3 {
		return e.ErrLongNationality
	}

	return nil
}
