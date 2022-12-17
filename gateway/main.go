package main

import (
	"encoding/json"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"server/src/csvs"
	"time"
)

func Cors() gin.HandlerFunc {
	c := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		//AllowHeaders:    []string{"Content-Type", "Access-Token", "Authorization"},
		AllowCredentials: true,
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		MaxAge:           6 * time.Hour,
	}

	return cors.New(c)
}

// Init 设置工作路径
func Init() {
	err := os.Chdir("/Users/biang/Desktop/github.com/BiangUncle/Guardian_Tales_Server")
	if err != nil {
		panic(err)
	}
	log.Infoln("设置工作路径")
}

func main() {

	Init()
	csvs.CheckLoadCsv()

	router := gin.Default()

	router.Use(Cors())

	router.POST("/summon", func(c *gin.Context) {
		typ := c.PostForm("type")
		log.Infoln(typ)

		heroList, err := ServerInstance().Summon(typ)
		if err != nil {
			log.Errorln(err)
			c.JSON(500, gin.H{
				"status": -1,
				"msg":    err.Error(),
				"data":   gin.H{},
			})
			return
		}

		b, err := json.Marshal(heroList)
		if err != nil {
			log.Errorln(err)
			c.JSON(500, gin.H{
				"status": -1,
				"msg":    err.Error(),
				"data":   gin.H{},
			})
			return
		}

		c.JSON(200, gin.H{
			"status": 0,
			"msg":    "ok",
			"data": gin.H{
				"items": string(b),
			},
		})
	})

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
