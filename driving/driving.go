package driving

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rifame/driven"
	"rifame/driving/handlers"
	"time"
)

type (
	Driving struct {
		Server *http.Server
	}

	ServerSettings struct {
		RunMode      string
		HttpPort     int
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}

	Settings struct {
		Server ServerSettings
	}
)

func (d Driving) Setup(driven driven.Driven, settings Settings) {
	gin.SetMode(settings.Server.RunMode)

	routes := initRoutes(driven)
	endPoint := fmt.Sprintf("127.0.0.1:%d", settings.Server.HttpPort)
	maxHeaderBytes := 1 << 20

	d.Server = &http.Server{
		Addr:           endPoint,
		Handler:        routes,
		ReadTimeout:    settings.Server.ReadTimeout,
		WriteTimeout:   settings.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	err := d.Server.ListenAndServe()
	if err != nil {
		log.Fatal("[fatal] failed to startup the server:", err)
	}
}

func initRoutes(driven driven.Driven) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api/categories")
	{
		api.GET("/a", handlers.FindAll(driven))
	}

	r.GET("/ping", func(c *gin.Context) {
		log.Println("ping handling")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
