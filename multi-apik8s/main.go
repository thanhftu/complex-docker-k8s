package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/thanhftu/api-multi/utils"
	"github.com/thanhftu/multi-apik8s/api"
	db "github.com/thanhftu/multi-apik8s/db/sqlc"
)

var (
	router = gin.Default()
)

func main() {
	// router.Use(cors.New(cors.Config{
	// 	AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
	// 	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowAllOrigins:  false,
	// 	AllowOriginFunc:  func(origin string) bool { return true },
	// 	MaxAge:           86400,
	// }))

	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3050"}
	// config.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}

	// router.GET("/api", controller.GetFibFromDB)
	// router.POST("/api/values", controller.GetWorkerResultHandler)
	// router.GET("/api/values/latest", controller.GetLatestFibHandler)
	// router.DELETE("/api/values/:id", controller.DeleteFibHandler)
	// router.GET("/api/allfib", controller.GetAllFinController)
	// router.Run(":8081")
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config file", err)
	}
	fmt.Println(config.DBDriver)
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddr)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
