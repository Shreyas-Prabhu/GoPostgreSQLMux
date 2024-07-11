package router

import (
	"encoding/json"
	"net/http"

	"github.com/Shreyas-Prabhu/GoPostgreSQLMux/controller"
	_ "github.com/Shreyas-Prabhu/GoPostgreSQLMux/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func CreateRouter() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Pong")
	})
	route.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	route.HandleFunc("/employee/createemployee", controller.AddEmployee).Methods("POST")
	route.HandleFunc("/employee/editemployee/{id}", controller.UpdateEmployee).Methods("PUT")
	route.HandleFunc("/employee/{id}", controller.GetEmployeeById).Methods("GET")
	route.HandleFunc("employee/deleteemployee/{id}", controller.DeleteEmployee).Methods("DELETE")
	route.HandleFunc("employee/employees", controller.GetAllEmployee).Methods("GET")

	route.HandleFunc("/asset/createasset", controller.AddAsset).Methods("POST")
	route.HandleFunc("/asset/editasset/{id}", controller.UpdateAsset).Methods("PUT")
	route.HandleFunc("/asset/{id}", controller.GetAssetById).Methods("GET")
	route.HandleFunc("/asset/deleteAsset/{id}", controller.DeleteAsset).Methods("DELETE")
	route.HandleFunc("/asset/assets", controller.GetAllAsset).Methods("GET")

	route.HandleFunc("/mapping/assignassetmapping", controller.AddMapping).Methods("POST")
	route.HandleFunc("/mapping/getallassets/{empid}", controller.GetAllMappingForEmp).Methods("GET")
	route.HandleFunc("/mapping/removeassetmapping/{id}", controller.RemoveMapping).Methods("DELETE")

	route.HandleFunc("/dashboard", controller.GetDashboard).Methods("GET")

	// employeeRouter := route.PathPrefix("/employee").Subrouter()
	// employeeRouter.HandleFunc("/createemployee", controller.AddEmployee).Methods("POST")
	// employeeRouter.HandleFunc("/{id}", controller.GetEmployeeById).Methods("GET")
	// employeeRouter.HandleFunc("/editemployee/{id}", controller.UpdateEmployee).Methods("PUT")
	// employeeRouter.HandleFunc("/deleteemployee/{id}", controller.DeleteEmployee).Methods("DELETE")
	// employeeRouter.HandleFunc("/employees", controller.GetAllEmployee).Methods("GET")

	// assetRouter := route.PathPrefix("/asset").Subrouter()
	// assetRouter.HandleFunc("/createasset", controller.AddAsset).Methods("POST")
	// assetRouter.HandleFunc("/{id}", controller.GetAssetById).Methods("GET")
	// assetRouter.HandleFunc("/editasset/{id}", controller.UpdateAsset).Methods("PUT")
	// assetRouter.HandleFunc("/deleteAsset/{id}", controller.DeleteAsset).Methods("DELETE")
	// assetRouter.HandleFunc("/assets", controller.GetAllAsset).Methods("GET")

	// mappingRouter := route.PathPrefix("/mapping").Subrouter()
	// mappingRouter.HandleFunc("/assignassetmapping", controller.AddMapping).Methods("POST")
	// mappingRouter.HandleFunc("/getallassets/{empid}", controller.GetAllMappingForEmp).Methods("GET")
	// mappingRouter.HandleFunc("/removeassetmapping/{id}", controller.RemoveMapping).Methods("DELETE")

	// dashboardRouter := route.PathPrefix("/dashboard").Subrouter()
	// dashboardRouter.HandleFunc("", controller.GetDashboard).Methods("GET")

	return route
}
