package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/thanhftu/multi-apik8s/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
	}
	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/fibs", server.CreateFibonacciHandler)
	router.GET("/lastcreated", server.GetLatestFibonacciHandler)
	router.GET("/getbyid/:id", server.GetFibonacciByIDHandler)
	router.GET("/fibnumbers", server.ListFibonacciHandler)
	router.DELETE("/fibnumbers/:id", server.DeleteFibHandler)
	server.router = router
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}
