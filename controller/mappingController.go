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

// @Title AssetMapping
// @Summary Create a new Asset Mapping
// @Description Create a new Asset Mapping with the input paylod
// @Tags Mapping
// @Accept  json
// @Produce  json
// @Param EmployeeAsset body model.EmployeeAsset true "Create Employee Asset Mapping"
// @Success 200
// @Router /mapping/assignassetmapping [post]
func AddMapping(w http.ResponseWriter, r *http.Request) {
	var mapping model.EmployeeAsset
	json.NewDecoder(r.Body).Decode(&mapping)
	mapping.Id = uuid.New()
	err := helper.AddMappingHelper(mapping)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode("Mapping created")
}

// @Title GetAllMappingForEmp
// @Summary get Asset Mapping for an employee id
// @Description gets all asset Mapping by providing path variable as emp id
// @Tags Mapping
// @Produce  json
// @Param empid path uuid.UUID true "asset ID"
// @Success 200
// @Router /mapping/getallassets/{empid} [get]
func GetAllMappingForEmp(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	empid, err := uuid.Parse(m["empid"])
	if err != nil {
		log.Println(err)
	}

	asset, err := helper.GetAllMappingWithEmpIdHelper(empid)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	response := make(map[string]interface{})
	response["Employee Id"] = empid
	response["Assets"] = asset
	json.NewEncoder(w).Encode(response)
}

// @Title RemoveMapping
// @Summary Remove asset mapping
// @DescriptionRemove asset mapping by providing path variable as mapping id
// @Tags Mapping
// @Produce  json
// @Param id path uuid.UUID true "mapping ID"
// @Success 200
// @Router /mapping/removeassetmapping/{id} [delete]
func RemoveMapping(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	mapid, err := uuid.Parse(m["id"])
	if err != nil {
		log.Println(err)
	}
	err = helper.DeleteMappingHelper(mapid)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode("Mapping Deleted")

}
