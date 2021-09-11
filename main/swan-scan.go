package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"
	"swap-scan/blockchain/browsersync/nbai"
	"swap-scan/blockchain/initclient/bscclient"
	"swap-scan/blockchain/initclient/nbaiclient"
	"swap-scan/blockchain/schedule"
	"swap-scan/common/constants"
	"swap-scan/config"
	"swap-scan/database"
	"swap-scan/logs"
	"swap-scan/routers"
	"time"
)

func main() {
	LoadEnv()

	// init database
	db := database.Init()

	initMethod()

	conf := config.GetConfig()
	fmt.Println(conf)

	go schedule.RedoMappingSchedule()

	go nbai.NbaiBlockBrowserSyncAndEventLogsSync()

	defer func() {
		err := db.Close()
		if err != nil {
			logs.GetLogger().Error(err)
		}
	}()

	r := gin.Default()
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	v1 := r.Group("/api/v1")
	routers.EventLogManager(v1.Group(constants.URL_EVENT_PREFIX))

	err := r.Run(":" + config.GetConfig().Port)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

}

func initMethod() string {
	config.InitConfig("")
	nbaiclient.ClientInit()
	bscclient.ClientInit()
	return ""
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logs.GetLogger().Error(err)
	}
}
