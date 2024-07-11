package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Shreyas-Prabhu/GoPostgreSQLMux/helper"
	"github.com/Shreyas-Prabhu/GoPostgreSQLMux/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// @Title CreateEmployee
// @Summary Create a new employee
// @Description Create a new employee with the input paylod
// @Tags Employee
// @Accept  json
// @Produce  json
// @Param employee body model.Employee true "Create Employee"
// @Success 200
// @Router /employee/createemployee [post]
func AddEmployee(w http.ResponseWriter, r *http.Request) {
	var emp model.Employee
	json.NewDecoder(r.Body).Decode(&emp)
	emp.Empid = uuid.New()
	err := helper.AddEmpHelper(emp)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode("Inserted")
}

// @Title UpdateEmployee
// @Summary Updates Employee
// @Description updates an employee with the input paylod
// @Tags Employee
// @Accept  json
// @Produce  json
// @Param employee body model.Employee true "Update Employee"
// @Success 200
// @Router /employee/editemployee/{id} [put]
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	empid, err := uuid.Parse(m["id"])
	if err != nil {
		log.Println(err)
	}
	var emp model.Employee
	json.NewDecoder(r.Body).Decode(&emp)
	err = helper.UpdateHelper(emp, empid)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode("Updated")
}

// @Title GetEmployeeById
// @Summary get Employee
// @Description gets an employee by providing path variable as emp id
// @Tags Employee
// @Produce  json
// @Param id path uuid.UUID true "Employee ID"
// @Success 200 {object} model.Employee
// @Router /employee/{id} [get]
func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	empid, err := uuid.Parse(m["id"])
	if err != nil {
		log.Println(err)
	}
	emp, err := helper.GetEmpByIdHelper(empid)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(emp)
}

// @Title DeleteEmployee
// @Summary Deletes Employee
// @Description delete an employee
// @Tags Employee
// @Produce  json
// @Param id path uuid.UUID true "Employee ID"
// @Success 200
// @Router /employee/deleteemployee/{id} [delete]
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	empid, err := uuid.Parse(m["id"])
	if err != nil {
		log.Println(err)
	}
	err = helper.DeleteHelper(empid)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode("Deleted")
}

// @Title GetAllEmployee
// @Summary gets all Employees
// @Description gets all Employees
// @Tags Employee
// @Produce  json
// @Success 200 {object} []model.Employee
// @Router /employee/employees [get]
func GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	emp, err := helper.GetAllHelper()
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(emp)
}
