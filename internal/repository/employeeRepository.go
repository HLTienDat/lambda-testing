package repository

import (
	"database/sql"
	"fmt"
	model "lambda-test/internal/repository/model"
	log "lambda-test/logs"
	"sync"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "091372"
	dbname   = "Users"
)

var (
	database *sql.DB
	once     sync.Once
)

type EmployeeRepository interface {
	ReadEmployee(id string) (model.Employee, error)
	ReadAllEmployee(id int) (model.Employee, error)
	CreateEmployee(e model.Employee) error
	Update(id int, e model.Employee) error
	DeleteEmployee(id int) error
}

func ConnectDB() (*sql.DB, error) {
	once.Do(func() {
		db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
		if err != nil {
			log.ErrorLog.Printf("Error connect to database: %s\n", err)
		}
		database = db
	})
	return database, nil
}

func ReadEmployee(id string) (model.Employee, error) {
	var e model.Employee
	db, err := ConnectDB()
	if err != nil {
		return e, err
	}
	row := db.QueryRow("SELECT id, name,dob, email,phone,citizenId,address FROM Employee WHERE id=$1", id)
	err = row.Scan(&e.Id, &e.Name, &e.Dob, &e.Email, &e.Phone, &e.CitizenId, &e.Address)
	if err != nil {
		log.ErrorLog.Printf("Error get Employee %v, %s\n", id, err)
	}
	return e, nil
}
func ReadAllEmployee() ([]model.Employee, error) {
	var employees []model.Employee
	db, err := ConnectDB()
	if err != nil {
		return employees, err
	}
	rows, err := db.Query("SELECT id, name,dob, email,phone,citizenId,address FROM Employee")
	if err != nil {
		log.ErrorLog.Printf("Error get Employees %s\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e model.Employee
		err := rows.Scan(&e.Id, &e.Name, &e.Dob, &e.Email, &e.Phone, &e.CitizenId, &e.Address)
		if err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}
	return employees, nil
}

func CreateEmployee(e model.Employee) error {
	var id int
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	err = db.QueryRow("INSERT INTO Employee(name, dob, email, phone, citizenId, address) VALUES($1, $2,$3,$4,$5,$6) RETURNING id", e.Name, e.Dob, e.Email, e.Phone, e.CitizenId, e.Address).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEmployee(id string, e model.Employee) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE Employee SET name=$1, dob=$2, email=$3, phone=$4, citizenId=$5, address=$6 WHERE id=$7", e.Name, e.Dob, e.Email, e.Phone, e.CitizenId, e.Address, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEmployee(id string) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM EMPLOYEE WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
