package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
	JWT struct {
		SecretKey   string
		ExpireHours int
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")                //设置配置文件名字
	viper.SetConfigType("yml")                   //设置文件类型
	viper.AddConfigPath("./config")              //设置文件目录
	if err := viper.ReadInConfig(); err != nil { //读取配置
		log.Fatalf("Error reading config file: %v", err)
	}
	AppConfig = &Config{} //声明一个结构体实例

	if err := viper.Unmarshal(AppConfig); err != nil { //把读取到的数据放入声明的结构体中
		log.Fatalf("Unable to decode into struct: %v", err)
	}
	fmt.Println("AppConfig", AppConfig.App.Port)

	fmt.Println("JwtConfig", AppConfig.JWT.ExpireHours)

	initDB() //初始化连接数据库
}
