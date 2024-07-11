package helper

import (
	"errors"
	"log"

	"github.com/Shreyas-Prabhu/GoPostgreSQLMux/database"
	"github.com/Shreyas-Prabhu/GoPostgreSQLMux/model"
	"github.com/google/uuid"
)

var DB = database.Connect()

func init() {
	log.Printf("Automigrating Employee schema...")
	DB.AutoMigrate(&model.Employee{})

}

func AddEmpHelper(emp model.Employee) error {
	res, err := GetAllHelper()
	if err != nil {
		return nil
	}
	for _, v := range res {
		if v.EmployeeEmail == emp.EmployeeEmail {
			return errors.New("employee already present")
		}
	}
	err = DB.Create(&emp).Error
	if err != nil {
		return err
	}
	return nil
}

func GetEmpByIdHelper(id uuid.UUID) (model.Employee, error) {
	var emp model.Employee
	err := DB.Find(&emp, id).Error
	if err != nil {
		return emp, err
	}
	return emp, nil

}

func UpdateHelper(emp model.Employee, id uuid.UUID) error {
	res, err := GetEmpByIdHelper(id)
	if err != nil {
		return err
	}
	err = DB.Model(&res).Updates(model.Employee{
		FirstName:        emp.FirstName,
		LastName:         emp.LastName,
		Gender:           emp.Gender,
		EmployeeEmail:    emp.EmployeeEmail,
		PhoneNumber:      emp.PhoneNumber,
		Address:          emp.Address,
		BloodGroup:       emp.BloodGroup,
		EmergencyContact: emp.EmergencyContact,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteHelper(id uuid.UUID) error {
	var emp model.Employee
	err := DB.Delete(&emp, id).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllHelper() ([]model.Employee, error) {
	var emp []model.Employee
	err := DB.Find(&emp).Error
	if err != nil {
		return nil, err
	}
	return emp, nil
}
