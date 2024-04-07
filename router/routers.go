package router

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
	// ad.HandleFunc("/all", controller.GetAllHandler).Methods("GET")
	ad.HandleFunc("", controller.CreateHandler).Methods("POST")
	// 路由設定api/v1/ad?offset=1&limit=3&age=26
	ad.HandleFunc("", controller.SearchAdsHandler).Methods("GET")

	return router
}
