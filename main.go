package main

import (
	"api/model"
	"api/router"
	"api/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {
	 //連線資料庫
	err := sql.InitSql()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	//AutoMigrate是GORM庫的一個方法，它將自動創建數據庫表，如果表已經存在，它將更新表的結構。
	sql.Connect.AutoMigrate(&model.Ad{}) 
	sql.Connect.AutoMigrate(&model.AdCondition{})

	r := router.SetRouter()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
