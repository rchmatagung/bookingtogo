package models

type CustomerRequest struct {
	NationalityId int64               `json:"nationality_id"`
	CustomerName  string              `json:"customer_name"`
	CustomerDOB   string              `json:"customer_dob"`
	CustomerPhone string              `json:"customer_phone"`
	CustomerEmail string              `json:"customer_email"`
	FamilyList    []FamilyListRequest `json:"family_list"`
}

type UpdateCustomerRequest struct {
	CustomerId    int64                 `json:"customer_id"`
	NationalityId int64                 `json:"nationality_id"`
	CustomerName  string                `json:"customer_name"`
	CustomerDOB   string                `json:"customer_dob"`
	CustomerPhone string                `json:"customer_phone"`
	CustomerEmail string                `json:"customer_email"`
	FamilyList    []UpdateFamilyRequest `json:"family_list"`
}

type UpdateFamilyRequest struct {
	FamilyListId       int64  `json:"family_list_id"`
	FamilyListRelation string `json:"family_list_relation"`
	FamilyListName     string `json:"family_list_name"`
	FamilyListDOB      string `json:"family_list_dob"`
	IsDelete           bool   `json:"is_delete"`
}

type FamilyListRequest struct {
	FamilyListRelation string `json:"family_list_relation"`
	FamilyListName     string `json:"family_list_name"`
	FamilyListDOB      string `json:"family_list_dob"`
}
