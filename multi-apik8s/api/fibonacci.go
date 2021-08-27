package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/thanhftu/multi-apik8s/db/sqlc"
	"github.com/thanhftu/multi-apik8s/service"
)

const (
	numberFibDisplay = 3
)

type requestData struct {
	Index string `json:"index"`
}

func (server *Server) CreateFibonacciHandler(c *gin.Context) {
	var reqData requestData
	if err := c.ShouldBindJSON(&reqData); err != nil {

		c.JSON(http.StatusBadRequest, errors.New("Bad request"))
		return
	}
	index, _ := strconv.ParseInt(reqData.Index, 10, 64)
	val, _ := service.GetFibFromRedisWorker(index)
	createFibonacciParam := db.CreateFibonacciParams{
		Index: index,
		Value: val,
	}
	fib, err := server.store.CreateFibonacci(c, createFibonacciParam)
	if err != nil {
		c.JSON(http.StatusNotImplemented, err.Error())
		return
	}
	fmt.Println("index: ", reqData.Index)
	fmt.Println("value: ", fib.Value)
	c.JSON(http.StatusOK, fib)
}

func (server *Server) GetLatestFibonacciHandler(c *gin.Context) {
	fibonacciNumner, err := server.store.GetLatestCreatedFibonacci(c)
	if err != nil {
		c.JSON(http.StatusNotImplemented, err.Error())
		return
	}
	c.JSON(http.StatusOK, fibonacciNumner)
}

func (server *Server) GetFibonacciByIDHandler(c *gin.Context) {
	id := c.Param("id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("id should be a number"))
		return
	}

	result, err := server.store.GetFibonacciByID(c, id64)
	if err != nil {
		c.JSON(http.StatusNotImplemented, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}

func (server *Server) ListFibonacciHandler(c *gin.Context) {

	results, err := server.store.ListFibonaccis(c, numberFibDisplay)
	if err != nil {
		c.JSON(http.StatusNotImplemented, err.Error())
		return
	}
	c.JSON(http.StatusOK, results)
}

func (server *Server) DeleteFibHandler(c *gin.Context) {
	id := c.Param("id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	if err := server.store.DeleteFibonacci(c, id64); err != nil {
		c.JSON(http.StatusNotImplemented, err)
	}
	c.JSON(http.StatusOK, id)
}
