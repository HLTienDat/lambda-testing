package usecase

import (
	entity "lambda-test/internal/entity"
	"lambda-test/internal/repository"
	model "lambda-test/internal/repository/model"
	log "lambda-test/logs"

	"github.com/jinzhu/copier"
)

type EmployeeUsecase interface {
	GetEmployees() ([]model.Employee, error)
	GetEmployee() ([]model.Employee, error)
	CreateEmployee(employee model.Employee) error
	UpdateEmployee(id int, employee model.Employee) error
	DeleteEmployee(id int) error
}

func GetEmployees(req entity.GetEmployeesRequest) (entity.GetEmployeesResponse, error) {
	resp := entity.GetEmployeesResponse{}
	log.InfoLog.Printf("Get Employees - Calling to Repository\n")
	employees, err := repository.ReadAllEmployee()
	if err != nil {
		log.ErrorLog.Printf("Get Employees - ERROR calling to Repository: %s\n", err)
		return resp, err
	}
	copier.Copy(&resp, &employees)
	return resp, nil
}

func GetEmployee(req entity.GetEmployeeRequest) (entity.GetEmployeeResponse, error) {
	resp := entity.GetEmployeeResponse{
		RequestID: req.RequestID,
	}
	log.InfoLog.Printf("Get Employee - Calling to Repository\n")
	employee, err := repository.ReadEmployee(req.EmployeeID)
	if err != nil {
		log.ErrorLog.Printf("Get Employee - ERROR calling to Repository: %s\n", err)
		return resp, err
	}

	copier.Copy(&resp, &employee)
	return resp, nil
}

func CreateEmployee(req entity.CreateEmployeeRequest) (entity.PostEmployeeResponse, error) {
	resp := entity.PostEmployeeResponse{
		RequestID: req.RequestID,
	}
	emp := model.Employee{}
	log.InfoLog.Printf("Create Employee - Calling to Repository\n")
	copier.Copy(&emp, &req.Data)
	err := repository.CreateEmployee(emp)
	if err != nil {
		log.ErrorLog.Printf("Create Employee - ERROR calling to Repository: %s\n", err)
		resp.Status = "Aborted"
		return resp, err
	}
	resp.Status = "Complete"

	return resp, nil
}

func UpdateEmployee(req entity.UpdateEmployeeRequest) (entity.PostEmployeeResponse, error) {
	resp := entity.PostEmployeeResponse{
		RequestID: req.RequestID,
	}
	emp := model.Employee{}
	copier.Copy(&emp, &req.Data)
	log.InfoLog.Printf("Update Employee - Calling to Repository\n")
	err := repository.UpdateEmployee(req.EmployeeID, emp)
	if err != nil {
		log.ErrorLog.Printf("Update Employee - ERROR calling to Repository: %s\n", err)
		resp.Status = "Aborted"
		return resp, err
	}
	resp.Status = "Complete"
	return resp, nil
}

func DeleteEmployee(req entity.DeleteEmployeeRequest) (entity.DeleteEmployeeResponse, error) {
	resp := entity.DeleteEmployeeResponse{
		RequestID: req.RequestID,
	}

	log.InfoLog.Printf("Delete Employee - Calling to Repository\n")
	err := repository.DeleteEmployee(req.EmployeeID)
	if err != nil {
		log.ErrorLog.Printf("Delete Employee - ERROR calling to Repository: %s\n", err)
		resp.Status = "Aborted"
		return resp, err
	}
	resp.Status = "Complete"
	return resp, nil
}
