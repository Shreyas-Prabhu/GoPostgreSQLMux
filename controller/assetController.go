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

// @Title CreateAsset
// @Summary Create a new Asset
// @Description Create a new Asset with the input paylod
// @Tags Asset
// @Accept  json
// @Produce  json
// @Param Asset body model.Asset true "Create Asset"
// @Success 200
// @Router /asset/createasset [post]
func AddAsset(w http.ResponseWriter, r *http.Request) {
	var asset model.Asset
	json.NewDecoder(r.Body).Decode(&asset)
	asset.AssetId = uuid.New()
	err := helper.AddAssetHelper(asset)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode("Asset Inserted")
}

// @Title UpdateAsset
// @Summary Updates Asset
// @Description updates an asset with the input paylod
// @Tags Asset
// @Accept  json
// @Produce  json
// @Param asset body model.Asset true "Update asset"
// @Success 200
// @Router /asset/editasset/{id} [put]
func UpdateAsset(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	assetid, err := uuid.Parse(m["id"])
	if err != nil {
		log.Println(err)
	}
	var asset model.Asset
	json.NewDecoder(r.Body).Decode(&asset)
	err = helper.UpdateAssetHelper(asset, assetid)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode("Asset Updated")
}

// @Title GetAssettById
// @Summary get Asset
// @Description gets an asset by providing path variable as asset id
// @Tags Asset
// @Produce  json
// @Param id path uuid.UUID true "asset ID"
// @Success 200 {object} model.Asset
// @Router /asset/{id} [get]
func GetAssetById(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	assetid, err := uuid.Parse(m["id"])
	if err != nil {
		log.Println(err)
	}
	asset, err := helper.GetAssetByIdHelper(assetid)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(asset)
}

// @Title Asset
// @Summary Deletes Asset
// @Description delete an asset
// @Tags Asset
// @Produce  json
// @Param id path uuid.UUID true "asset ID"
// @Success 200
// @Router /deleteasset/{id} [delete]
func DeleteAsset(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	assetid, err := uuid.Parse(m["id"])
	if err != nil {
		log.Println(err)
	}
	err = helper.DeleteAssetHelper(assetid)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode("Asset Deleted")
}

// @Title GetAllAsset
// @Summary gets all assets
// @Description gets all assets
// @Tags Asset
// @Produce  json
// @Success 200 {object} []model.Asset
// @Router /asset/assets [get]
func GetAllAsset(w http.ResponseWriter, r *http.Request) {
	asset, err := helper.GetAllAssetHelper()
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(asset)
}
