package routes

import (
	"api/controller"

	"github.com/gorilla/mux"
)

func SetRouter() *mux.Router {
	router := mux.NewRouter()
	
	// 定義共同的 URL 前綴
    apiV1 := router.PathPrefix("/api/v1").Subrouter()
    ad := apiV1.PathPrefix("/ad").Subrouter()

    // 定義路由
    ad.HandleFunc("/", controller.GetAll).Methods("GET")
    ad.HandleFunc("/", controller.Create).Methods("POST")

    return router
}
