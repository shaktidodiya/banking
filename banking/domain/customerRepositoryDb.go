package domain

import (
	"Desktop/golang/banking/errs"
	"Desktop/golang/banking/logger"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError){

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		logger.Error("Error while querying customer table "+ err.Error())
		return nil, errs.NewUnexpectedError("unexpected DB error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
		if err != nil {
			logger.Error("Error while querying customer table "+ err.Error())
			return nil, errs.NewUnexpectedError("unexpected DB error")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) GetById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected DB error")
	}
	return &c, nil
}

func (d CustomerRepositoryDb) GetByStatus(status string) ([]Customer, *errs.AppError) {
	//var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)

	fmt.Println(status)
	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql, status)
		//rows, err = d.client.Query(findAllSql, status)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
		//rows, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table "+ err.Error())
		return nil, errs.NewUnexpectedError("unexpected DB error")
	}

	//err = sqlx.StructScan(rows, &customers)
	//if err != nil {
	//	logger.Error("Error while querying customer table "+ err.Error())
	//	return nil, errs.NewUnexpectedError("unexpected DB error")
	//}

	//for rows.Next() {
	//	var c Customer
	//	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	//	if err != nil {
	//		logger.Error("Error while querying customer table "+ err.Error())
	//		return nil, errs.NewUnexpectedError("unexpected DB error")
	//	}
	//	customers = append(customers, c)
	//}
	return customers, nil
}

func NewCustomerRepositoryDb(dbCleint *sqlx.DB) CustomerRepositoryDb{
	return CustomerRepositoryDb{dbCleint}
}
