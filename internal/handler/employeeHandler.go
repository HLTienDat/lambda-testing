package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	entity "lambda-test/internal/entity"
	usecase "lambda-test/internal/usecase"
	log "lambda-test/logs"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler interface {
	GetEmployees(c *gin.Context)
}

func GetEmployees(c *gin.Context) {
	body := entity.GetEmployeesRequest{}
	if err := c.BindJSON(&body); err != nil {
		log.ErrorLog.Printf("Get Employees - ERROR Reading Body: %s\n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	resp, err := usecase.GetEmployees(body)
	if err != nil {
		log.ErrorLog.Printf("Get Employees - ERROR calling to Usecase: %s\n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Cannot get Employees: %s", err)})
		return
	}
	c.IndentedJSON(http.StatusOK, resp)
}
func GetEmployee(c *gin.Context) {
	body := entity.GetEmployeeRequest{}
	if err := c.BindJSON(&body); err != nil {
		log.ErrorLog.Printf("Get Employee - ERROR Reading Body: %s\n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	resp, err := usecase.GetEmployee(body)
	if err != nil {
		if err == sql.ErrNoRows {
			log.WarningLog.Printf("Get Employee - ERROR calling to Usecase: %s\n", err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Employee not found"})
			return
		}
		log.ErrorLog.Printf("Get Employee - ERROR calling to Usecase: %s\n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Cannot get employee ID: %v %s", body.EmployeeID, err)})
		return
	}
	c.IndentedJSON(http.StatusOK, resp)
}
func CreateEmployee(c *gin.Context) {
	body := entity.CreateEmployeeRequest{}
	if err := c.BindJSON(&body); err != nil {
		log.ErrorLog.Printf("Create Employee - ERROR Reading Body: %s\n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	resp, err := usecase.CreateEmployee(body)
	if err != nil {
		log.ErrorLog.Printf("Create Employee - ERROR calling to Usecase: %s\n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Cannot create new employee: %s", err)})
		return
	}
	c.IndentedJSON(http.StatusCreated, resp)
}

func UpdateUser(c *gin.Context) {
	body := entity.UpdateEmployeeRequest{}
	if err := c.BindJSON(&body); err != nil {
		log.ErrorLog.Printf("Update Employee - ERROR Reading Body: %s\n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	resp, err := usecase.UpdateEmployee(body)
	if err != nil {
		log.ErrorLog.Printf("Update Employee - ERROR calling to Usecase: %s\n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Cannot update employee ID: %v %s", body.EmployeeID, err)})
		return
	}
	c.IndentedJSON(http.StatusCreated, resp)
}

func DelUsers(c *gin.Context) {
	body := entity.DeleteEmployeeRequest{}
	if err := c.BindJSON(&body); err != nil {
		log.ErrorLog.Printf("Delete Employee - ERROR Reading Body: %s\n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	resp, err := usecase.DeleteEmployee(body)
	if err != nil {
		log.ErrorLog.Printf("Delete Employee - ERROR calling to Usecase: %s\n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Cannot delete employee: %s", err)})
		return
	}
	c.IndentedJSON(http.StatusCreated, resp)
}
