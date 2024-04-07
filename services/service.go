package services

import (
	"api/model"
	"api/sql"
	"fmt"
	"time"
)

// func GetALlAd() (ads []*model.Ad, err error) {
// 	err = sql.Connect.Preload("Conditions").Find(&ads).Error //Preload是GORM庫的一個方法，它將預加載關聯的數據
// 	// Error是GORM庫的一個方法，它將返回一個錯誤，如果有錯誤，我們將在控制器中處理這個錯誤。
// 	// sql.Connect.Find(&ads)是GORM庫的一個方法，它將從數據庫中獲取所有的廣告，並將它們存儲在一個切片中。
// 	return
// }

func CreateAd(ad *model.Ad) (err error) {
	// 新增 Ad，因為有關連的 AdCondition，所以一併新增
	// sql.Connect.Create(ad)是GORM庫的一個方法，它將創建一個新的廣告，並將其存儲在數據庫中。

	if err = sql.Connect.Create(ad).Error; err != nil {
		return
	}

	return
}

func SearchAds(gender, country, platform string, age, offset, limit int) (res *model.Response, err error) {
	var results []model.Result

	// 建立 DB 查詢
	db := sql.Connect.Model(&model.Ad{})

	// 查詢活躍的廣告
	now := time.Now()
	db = db.Where("start_at < ? AND end_at > ?", now, now)

	// 連接 Ad 和 AdCondition 表
	db = db.Joins("JOIN ad_conditions ON ad_conditions.ad_id = ads.id")

	// 查詢國家條件
	if country != "" {
		// db = db.Where("JSON_CONTAINS(ad_conditions.country, ?)", fmt.Sprintf("\"%s\"", country))
		db = db.Where("JSON_CONTAINS(ad_conditions.country, ?) OR JSON_CONTAINS(ad_conditions.country, '\"ALL\"')", fmt.Sprintf("\"%s\"", country))
	}

	if platform != "" {
		db = db.Where("JSON_CONTAINS(ad_conditions.platform, ?)", fmt.Sprintf("\"%s\"", platform))
	}

	// 查詢性別條件
	if gender != "" {
		db = db.Where("JSON_CONTAINS(ad_conditions.gender, ?)", fmt.Sprintf("\"%s\"", gender))
	}

	// 先排序再分頁
	db = db.Select("title, end_at").Order("end_at asc")

	// 執行查詢時，只取 offset ~ offset+limit 的資料
	err = db.Offset(offset).Limit(limit).Find(&results).Error

	// 將結果包裝在 "items" 物件中
	response := model.Response{
		Items: results,
	}

	return &response, err
}
