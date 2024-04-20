package models

type Customer struct {
	CustomerId    int64        `json:"customer_id" db:"cst_id"`
	Nationality   string       `json:"nationality" db:"nationality"`
	CustomerName  string       `json:"customer_name" db:"cst_name"`
	CustomerDOB   string       `json:"customer_dob" db:"cst_dob"`
	CustomerPhone string       `json:"customer_phone" db:"cst_phoneNum"`
	CustomerEmail string       `json:"customer_email" db:"cst_email"`
	FamilyList    []FamilyList `json:"family_list" gorm:"-"`
}

type CustomerList struct {
	CustomerId    int64  `json:"customer_id" db:"cst_id"`
	CustomerName  string `json:"customer_name" db:"cst_name"`
	CustomerDOB   string `json:"customer_dob" db:"cst_dob"`
	CustomerPhone string `json:"customer_phone" db:"cst_phoneNum"`
	CustomerEmail string `json:"customer_email" db:"cst_email"`
}

type FamilyList struct {
	FamilyListId       int64  `json:"family_list_id" db:"fl_id"`
	CustomerId         int64  `json:"customer_id" db:"cst_id"`
	FamilyListRelation string `json:"family_list_relation" db:"fl_relation"`
	FamilyListName     string `json:"family_list_name" db:"fl_name"`
	FamilyListDOB      string `json:"family_list_dob" db:"fl_dob"`
}
