package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/satriyoaji/echo-restful-api/database"
	"github.com/satriyoaji/echo-restful-api/helpers"
	"net/http"
)

type Employee struct {
	Id      int    `json:"id"`
	Name    string `json:"nama" validate:"required"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone"`
}

func FetchAllEmployees() (Response, error) {
	var obj Employee
	var arrayObj []Employee
	var res Response

	con := database.GetConnection()

	sqlStatement := "SELECT * FROM employees"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	helpers.OutputPanicError(err)

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Address, &obj.Phone)
		if err != nil {
			return res, err
		}

		arrayObj = append(arrayObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrayObj

	return res, nil
}

func StoreEmployee(name, address, phone string) (Response, error) {
	var res Response

	v := validator.New()
	employeeStruct := Employee{
		Name:    name,
		Address: address,
		Phone:   phone,
	}
	// validation input
	err := v.Struct(employeeStruct)
	if err != nil {
		return res, err
	}

	con := database.GetConnection()

	sqlStatement := "INSERT employees (name, address, phone) VALUES (?,?,?)"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(name, address, phone)
	if err != nil {
		return res, err
	}
	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully created"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdateEmployee(id int, name, address, phone string) (Response, error) {
	var res Response

	con := database.GetConnection()

	sqlStatement := "UPDATE employees set name = ?, address = ?, phone = ? WHERE id = ?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(name, address, phone, id)
	if err != nil {
		return res, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully updated"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteEmployee(id int) (Response, error) {
	var res Response

	con := database.GetConnection()

	sqlStatement := "DELETE FROM employees WHERE id = ?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully deleted"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
