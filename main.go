package main

import (
	"Go-Web-Study/config"
	"Go-Web-Study/db"
	"Go-Web-Study/web"
	"fmt"
	"log"
)

func main() {

	appConfig, err := config.LoadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.Init(appConfig.DataSource.Username, appConfig.DataSource.Password, appConfig.DataSource.Address, appConfig.DataSource.DBName)
	if err != nil {
		log.Printf("DB Init failed, cause: %s", err.Error())
	}
	web.Service(appConfig.Http.Service, appConfig.Http.Port)

	//
	//port := &appConfig.Http.PortΩ
	//
	//log.Printf("port: %s", *port)
	//
	//r := gin.New()
	//r.GET("/JSONP", func(c *gin.Context) {
	//	data := map[string]interface{}{
	//		"foo": "bar",
	//	}
	//
	//	// /JSONP?callback=x
	//	// 将输出：x({\"foo\":\"bar\"})
	//	c.JSONP(http.StatusOK, data)
	//})
	//
	//// 监听并在 0.0.0.0:8080 上启动服务
	//r.Run(":8080")

}
