package models

type CustomerInsertResponse struct {
	CustomerId    int64        `json:"customer_id" db:"cst_id"`
	NationalityId int64        `json:"nationality_id" db:"nationality_id"`
	CustomerName  string       `json:"customer_name" db:"cst_name"`
	CustomerDOB   string       `json:"customer_dob" db:"cst_dob"`
	CustomerPhone string       `json:"customer_phone" db:"cst_phoneNum"`
	CustomerEmail string       `json:"customer_email" db:"cst_email"`
	FamilyList    []FamilyList `json:"family_list" gorm:"-"`
}
