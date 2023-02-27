package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"

	"vue-api/api"
)

func main() {

	var cfg api.Configuration
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		log.Fatalf("Server init failed, err: %+v\n", err)
	}

	// logFile, err := os.OpenFile(cfg.LogFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("Open log file failed, err: %+v\n", err)
	// }
	// defer logFile.Close()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
	// log.SetOutput(logFile)
}
