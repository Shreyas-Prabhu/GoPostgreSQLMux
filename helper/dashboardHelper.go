package helper

import "github.com/google/uuid"

func GetDashboardHelper() (interface{}, error) {
	type Dashboard struct {
		Empid            uuid.UUID
		FirstName        string
		LastName         string
		Gender           string
		PhoneNumber      int64
		EmployeeEmail    string
		Address          string
		BloodGroup       string
		EmergencyContact int64
		AssetCount       int
	}

	var data []Dashboard
	var count int
	emp, err := GetAllHelper()
	if err != nil {
		return nil, err
	}
	for _, v := range emp {
		asset, err := GetAllMappingWithEmpIdHelper(v.Empid)
		if err != nil {
			return nil, err
		}
		count = len(asset)
		singleDashboard := Dashboard{
			Empid:            v.Empid,
			FirstName:        v.FirstName,
			LastName:         v.LastName,
			Gender:           v.Gender,
			PhoneNumber:      v.PhoneNumber,
			EmployeeEmail:    v.EmployeeEmail,
			Address:          v.Address,
			BloodGroup:       v.BloodGroup,
			EmergencyContact: v.EmergencyContact,
			AssetCount:       count,
		}
		data = append(data, singleDashboard)
	}
	return data, nil

}
