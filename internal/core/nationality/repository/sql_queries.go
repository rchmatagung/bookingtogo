package repository

const (
	InsertNationality = `INSERT INTO nationality (nationality_name, nationality_code) VALUES (?, ?) RETURNING *`

	SelectAllNationality = `SELECT * FROM nationality`
)
