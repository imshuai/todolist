package main

import (
	"log"
	"net"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
	err := config.Read("config.json")
	if err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}
	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	err = dbInit()
	if err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}
}

func main() {
	defer db.Close()

	e := gin.Default()

	e.Use(gin.BasicAuth(gin.Accounts{
		"admin": "laoshuai01.*",
	}))

	e.LoadHTMLGlob("tpl/*")

	e.GET("/", handleIndex)
	e.GET("/missions", handleMissions)
	e.GET("/mission/:id", handleMission)
	e.POST("/mission", handleNewMission)
	e.DELETE("/mission/:id", handleDeleteMission)
	e.PATCH("/mission/:id", handleChangeMission)

	e.Run(net.JoinHostPort(config.ListenAddress, strconv.Itoa(config.ListenPort)))
	return
}
