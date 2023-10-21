package models

type People struct {
	ID          string `json:"id"                   db:"id"`
	Name        string `json:"name"                 db:"name"`
	Surname     string `json:"surname"              db:"surname"`
	Patronymic  string `json:"patronymic,omitempty" db:"patronymic"`
	Age         int    `json:"age"                  db:"age"`
	Gender      string `json:"gender"               db:"gender"`
	Nationality string `json:"nationality"          db:"nationality"`
}

type GetPeopleRequest struct {
	Name        string `json:"name,omitempty"`
	Surname     string `json:"surname,omitempty"`
	Patronymic  string `json:"patronymic,omitempty"`
	StartAge    int    `json:"start_age,omitempty"`
	EndAge      int    `json:"end_age,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Nationality string `json:"nationality,omitempty"`
	Limit       uint64 `json:"limit"`
	Offset      uint64 `json:"offset"`
}

type AbbreviatedPeopleRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}
