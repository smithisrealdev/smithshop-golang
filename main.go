package main

import (
	"github.com/gin-gonic/gin"
)

func init() {
	// env := os.Args[1]

	// API Start
	// gin.Logger()
	// log := gin.New()
	// log.Use(gin.Logger())

	// log.Info("", "Server start running on %s environment configuration", env)

	// config.SetConfigFile("config/" + env + ".yml")
	// if err := config.ReadInConfig(); err != nil {
	// 	log.Error("", "Fatal error env config (file): %s", err)
	// }
	// logger.SetSuffix()

	// getAppConfig()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
