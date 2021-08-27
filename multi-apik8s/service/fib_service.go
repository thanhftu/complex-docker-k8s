package service

import (
	"errors"
	"fmt"
	"os"

	"github.com/thanhftu/worker"
)

var RedisSource string

func init() {
	RedisSource = "localhost:6379"
	if v := os.Getenv("REDIS_ADDR"); v != "" {
		RedisSource = v
	}

	fmt.Println("Redis url: ", RedisSource)

}

func GetFibFromRedisWorker(index int64) (int64, error) {

	val, err := worker.WorkerRedisFib(index, RedisSource)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	return val, nil
}

// func GetFib(index string) (*fib.FibNumber, error) {
// 	index64, _ := strconv.ParseInt(index, 10, 64)
// 	fibnumber := &fib.FibNumber{
// 		Index: index64,
// 	}
// 	if err := fibnumber.GET(); err != nil {
// 		return nil, err
// 	}
// 	return fibnumber, nil
// }
// func GetLatest() (*fib.FibNumber, error) {
// 	var fiblastest fib.FibNumber
// 	if err := fiblastest.GETLATEST(); err != nil {
// 		return nil, err
// 	}
// 	return &fiblastest, nil
// }
// func GetAllFib() ([]fib.FibNumber, error) {
// 	return fib.GETALL()
// }

// func DeleteFib(ID string) error {
// 	ID64, _ := strconv.ParseInt(ID, 10, 64)
// 	fibnumber := &fib.FibNumber{
// 		ID: ID64,
// 	}
// 	if err := fibnumber.DELETE(); err != nil {
// 		return err
// 	}
// 	return nil
// }
