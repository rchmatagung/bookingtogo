package repository

const (
	InsertCustomer = `
		INSERT INTO customer (nationality_id, cst_name, cst_dob, cst_phoneNum, cst_email) 
		VALUES (?, ?, ?, ?, ?) 
		RETURNING cst_id as customer_id, cst_name as customer_name, cst_dob as customer_dob, cst_phoneNum as customer_phone, cst_email as customer_email, nationality_id
	`

	InsertFamilyList = `
		INSERT INTO family_list (cst_id, fl_relation, fl_name, fl_dob)
		VALUES (?, ?, ?, ?)
		RETURNING fl_id as family_list_id, cst_id as customer_id, fl_relation as family_list_relation, fl_name as family_list_name, fl_dob as family_list_dob
	`

	SelectAllCustomer = `
		SELECT cst_id as customer_id, cst_name as customer_name, TO_CHAR(cst_dob, 'YYYY-MM-DD') as customer_dob, cst_phoneNum as customer_phone, cst_email as customer_email
		from customer c 
	`

	SelectCustomerById = `
		SELECT 
			c.cst_id as customer_id,
			c.cst_name as customer_name,
			TO_CHAR(c.cst_dob, 'YYYY-MM-DD') as customer_dob,
			c.cst_phoneNum as customer_phone,
			c.cst_email as customer_email,
			CONCAT(n.nationality_name, ' (', n.nationality_code, ')') AS nationality
		from customer c
		inner join nationality n on c.nationality_id = n.nationality_id
		where c.cst_id = ?
	`

	SelectFamilyListByCustomerId = `
		SELECT 
			fl_id as family_list_id,
			cst_id as customer_id,
			fl_relation as family_list_relation,
			fl_name as family_list_name,
			fl_dob as family_list_dob 
		from family_list WHERE cst_id = ?
	`

	UpdateCustomer = `
		UPDATE customer
		SET nationality_id = ?, cst_name = ?, cst_dob = ?, cst_phoneNum = ?, cst_email = ?
		WHERE cst_id = ?
		RETURNING cst_id as customer_id, cst_name as customer_name, cst_dob as customer_dob, cst_phoneNum as customer_phone, cst_email as customer_email, nationality_id
	`

	UpdateFamilyList = `
		UPDATE family_list
		SET fl_relation = ?, fl_name = ?, fl_dob = ?
		WHERE fl_id = ?
		RETURNING fl_id as family_list_id, cst_id as customer_id, fl_relation as family_list_relation, fl_name as family_list_name, fl_dob as family_list_dob
	`

	DeleteCustomer = `
		DELETE FROM customer WHERE cst_id = ?
	`

	DeleteFamilyList = `
		DELETE FROM family_list WHERE fl_id = ?
	`
)