package sql

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connect *gorm.DB // Connect是一個指向gorm.DB的指針，它將用於與數據庫進行交互。

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

func (c *Config) getConf() *Config {
	// 這個函數將讀取config.yaml文件並將其解析為結構體，(c *Config)是一個接收器，它是一個指向Config結構體的指針，這意味著我們可以在函數中修改Config結構體的值。
	// 讀取YAML文件
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
	}

	// 將YAML解析為結構體
	err = yaml.Unmarshal(data, c)
	if err != nil {
		fmt.Println("Error parsing YAML:", err)
	}

	return c
}

func initDB() (err error) {
	var c Config

	//獲取yaml配置引數
	conf := c.getConf()

	//將yaml配置引數拼接成連線資料庫的url
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Name,
	)

	//連線資料庫
	Connect, err = gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{})
	return
}
