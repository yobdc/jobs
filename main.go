package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	_ "github.com/yobdc/jobs/conf"
	"github.com/yobdc/jobs/models"
	"log"
	"net/http"
)

var ti1 *models.TaskInstance

func main() {
	log.Println("hello")
	server := gin.New()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	server.GET("/", func(c *gin.Context) {

		t1 := models.NewTask("task1", "i am task1", "ping -c 1 baidu.com")
		t2 := models.NewTask("task2", "i am task2", "ping -c 30 baidu.com")
		t3 := models.NewTask("task3", "i am task3", "exit 2")
		t4 := models.NewTask("task4", "i am task4", "ping -c 2 baidu.com")
		t1.AddChild(t2)
		t1.AddChild(t3)
		t1.AddChild(t4)
		t2.AddChild(t4)
		t3.AddChild(t4)
		ti1 = t1.NewInstance()
		ti1.Start()

		c.JSON(http.StatusOK, gin.H{
			"message": ti1.ListInstances(),
		})
	})

	ginpprof.Wrapper(server)

	server.Run(":8080")
}
