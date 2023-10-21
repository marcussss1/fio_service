package e

import "errors"

var (
	// ErrLongName .
	ErrLongName = errors.New("имя человека не должно быть длиннее 50 символов")
	// ErrLongSurname .
	ErrLongSurname = errors.New("фамилия человека не должна быть длиннее 50 символов")
	// ErrLongPatronymic .
	ErrLongPatronymic = errors.New("отчество человека не должно быть длиннее 50 символов")
	// ErrInvalidAge .
	ErrInvalidAge = errors.New("возраст человека должен быть от 0 до 150 лет")
	// ErrInvalidGender .
	ErrInvalidGender = errors.New("пол человека должен быть male или female")
	// ErrLongNationality .
	ErrLongNationality = errors.New("национальность человека не должна быть длиннее 3 символов")
	// ErrInvalidPeopleID .
	ErrInvalidPeopleID = errors.New("id человека должен быть валидный UUID")
	// ErrPeopleNotFound .
	ErrPeopleNotFound = errors.New("человек не найден")
)
