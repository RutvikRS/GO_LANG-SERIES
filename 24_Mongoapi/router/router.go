package router

import (
	"github.com/gorilla/mux"
	"github.com/hiteshchoudhary/mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/TODOLISTS", controller.GetMyAllData).Methods("GET")
	router.HandleFunc("/api/data", controller.CreateTODO).Methods("POST")
	router.HandleFunc("/api/data/{id}", controller.MarkAsClicked).Methods("PUT")
	router.HandleFunc("/api/data/{id}", controller.DeleteTODOList).Methods("DELETE")
	router.HandleFunc("/api/deleteTODOLISTS", controller.DeleteAllTODOList).Methods("DELETE")

	return router
}
