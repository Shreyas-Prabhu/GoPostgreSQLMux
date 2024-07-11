package model

import "github.com/google/uuid"

// Employee represents the model for an employee
type Employee struct {
	Empid            uuid.UUID `json:"empid,omitempty" gorm:"primaryKey"`
	FirstName        string    `json:"firstname,omitempty" example:"Shreyas"`
	LastName         string    `json:"lastname,omitempty" example:"Prabhu"`
	Gender           string    `json:"gender,omitempty" example:"male"`
	PhoneNumber      int64     `json:"phonenumber,omitempty" example:"8787878787"`
	EmployeeEmail    string    `json:"employeeemail,omitempty" example:"ppp@gmail.com"`
	Address          string    `json:"address,omitempty" example:"mumbai"`
	BloodGroup       string    `json:"bloodgroup,omitempty" example:"AB-ve"`
	EmergencyContact int64     `json:"emergencycontact,omitempty" example:"84984894"`
}

type Asset struct {
	AssetId   uuid.UUID `json:"assetid,omitempty" gorm:"primaryKey"`
	AssetName string    `json:"assetname,omitempty" example:"Laptop"`
	AssetType string    `json:"assettype,omitempty" example:"Electronic"`
}

type EmployeeAsset struct {
	Id      uuid.UUID `json:"id,omitempty" gorm:"primaryKey"`
	EmpId   uuid.UUID `json:"empid,omitempty" example:"<Replace with the emp id>"`
	AssetId uuid.UUID `json:"assetid,omitempty" example:"<Replace with the asset id>"`
}
