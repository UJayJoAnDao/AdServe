package services

import (
	"api/model"
	"api/sql"
)

func GetALlAd() (ads []*model.Ad, err error) {
	err = sql.Connect.Find(&ads).Error
	// Error是GORM庫的一個方法，它將返回一個錯誤，如果有錯誤，我們將在控制器中處理這個錯誤。
	// sql.Connect.Find(&ads)是GORM庫的一個方法，它將從數據庫中獲取所有的廣告，並將它們存儲在一個切片中。
	return
}

func CreateAd(ad *model.Ad) (err error) {
	err = sql.Connect.Create(ad).Error
	// sql.Connect.Create(ad)是GORM庫的一個方法，它將創建一個新的廣告並將其保存到數據庫中。
	return
}

func SearchAd(condition *model.AdCondition) (ads []*model.Ad, err error) {
	err = sql.Connect.Where(condition).Find(&ads).Error
	// sql.Connect.Where(condition).Find(&ads)是GORM庫的一個方法，它將從數據庫中獲取符合條件的廣告，並將它們存儲在一個切片中。
	return
}
