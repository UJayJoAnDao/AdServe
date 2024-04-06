package controller

import (
	"api/model"
	"api/services"
	"encoding/json"
	"net/http"
)

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ads, err := services.GetALlAd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(ads)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var ad model.Ad
	json.NewDecoder(r.Body).Decode(&ad)
	err := services.CreateAd(&ad)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //http.Error是一個函數，它將返回一個錯誤響應，並將錯誤消息設置為響應的正文。
		return
	}
	json.NewEncoder(w).Encode(ad) //json.NewEncoder是一個函數，它將一個結構體編碼為JSON格式，並將其寫入一個io.Writer接口。
}
