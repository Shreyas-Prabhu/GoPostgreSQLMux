package helper

import (
	"log"

	"github.com/Shreyas-Prabhu/GoPostgreSQLMux/model"
	"github.com/google/uuid"
)

func init() {
	log.Printf("Automigrating Mapping schema...")
	DB.AutoMigrate(&model.EmployeeAsset{})

}

func AddMappingHelper(mapping model.EmployeeAsset) error {
	err := DB.Create(&mapping).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllMappingWithEmpIdHelper(empid uuid.UUID) ([]model.Asset, error) {
	var mapping []model.EmployeeAsset
	err := DB.Find(&mapping).Error
	if err != nil {
		return nil, err
	}
	var asset []model.Asset
	for _, v := range mapping {
		if v.EmpId == empid {
			res, err := GetAssetByIdHelper(v.AssetId)
			if err != nil {
				return nil, err
			}
			asset = append(asset, res)
		}
	}
	return asset, nil

}

func DeleteMappingHelper(mapid uuid.UUID) error {
	var mapping model.EmployeeAsset
	err := DB.Delete(&mapping, mapid).Error
	if err != nil {
		return err
	}
	return nil
}
