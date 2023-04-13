package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

type Config struct {
	Port        int    `mapstructure:"port"`
	Version     string `mapstructure:"version"`
	Mysqlconfig `mapstructure:"mysql"`
}
type Mysqlconfig struct {
	Port   int    `mapstructure:"port"`
	Ip     string `mapstructure:"ip"`
	User   string `mapstructure:"user"`
	Passwd string `mapstructure:"passwd"`
	Dbname string `mapstructure:"dbname"`
}

func main() {
	//设置默认值
	viper.SetDefault("fileDir", "./conf/")
	//读取配置文件
	viper.SetConfigFile("config.yaml") // 指定配置文件路径
	viper.SetConfigName("config")      // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")        // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("./conf/")     // 查找配置文件所在的路径

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
	//r := gin.Default()
	//r.GET("/version", func(c *gin.Context) {
	//	c.String(http.StatusOK, viper.GetString("version"))
	//})
	//r.Run()
	//从配置文件中反序列化
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("c:%#v\n", c)
}
