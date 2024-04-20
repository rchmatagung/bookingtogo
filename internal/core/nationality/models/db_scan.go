package models

type Nationality struct {
	NationalityId   int64  `json:"nationality_id"`
	NationalityName string `json:"nationality_name"`
	NationalityCode string `json:"nationality_code"`
}
