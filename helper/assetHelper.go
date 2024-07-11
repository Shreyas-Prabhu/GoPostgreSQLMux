package helper

import (
	"errors"
	"log"

	"github.com/Shreyas-Prabhu/GoPostgreSQLMux/model"
	"github.com/google/uuid"
)

func init() {
	log.Printf("Automigrating Asset schema...")
	DB.AutoMigrate(&model.Asset{})

}

func AddAssetHelper(asset model.Asset) error {
	res, err := GetAllAssetHelper()
	if err != nil {
		return nil
	}
	for _, v := range res {
		if v.AssetName == asset.AssetName {
			return errors.New("asset already present")
		}
	}
	err = DB.Create(&asset).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAssetByIdHelper(id uuid.UUID) (model.Asset, error) {
	var asset model.Asset
	err := DB.Find(&asset, id).Error
	if err != nil {
		return asset, err
	}
	return asset, nil

}

func UpdateAssetHelper(asset model.Asset, id uuid.UUID) error {
	res, err := GetAssetByIdHelper(id)
	if err != nil {
		return err
	}
	err = DB.Model(&res).Updates(model.Asset{
		AssetId:   asset.AssetId,
		AssetName: asset.AssetName,
		AssetType: asset.AssetType,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteAssetHelper(id uuid.UUID) error {
	var asset model.Asset
	err := DB.Delete(&asset, id).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllAssetHelper() ([]model.Asset, error) {
	var asset []model.Asset
	err := DB.Find(&asset).Error
	if err != nil {
		return nil, err
	}
	return asset, nil
}
