package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Shreyas-Prabhu/GoPostgreSQLMux/helper"
)

// @Title GetDashboard
// @Summary gets all data
// @Description gets all data of employee along with number of asset
// @Tags Dashboard
// @Produce  json
// @Success 200 {object} interface{}
// @Router /dashboard [get]
func GetDashboard(w http.ResponseWriter, r *http.Request) {
	data, err := helper.GetDashboardHelper()
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	response := make(map[string]interface{})
	response["assetList"] = data
	json.NewEncoder(w).Encode(response)
}
