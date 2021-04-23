package main

import (
	"fmt"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/hewenyu/gin-simple/initialize"
	"github.com/hewenyu/gin-simple/logger"
)

func main() {
	Router := initialize.Routers()

	endless.DefaultReadTimeOut = 10 * time.Millisecond
	endless.DefaultWriteTimeOut = 10 * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20

	endPoint := fmt.Sprintf(":%d", 8080)

	server := endless.NewServer(endPoint, Router)

	server.BeforeBegin = func(add string) {
		logger.Info(fmt.Sprintf("Actual pid is %d", syscall.Getpid()))
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Error(fmt.Sprintf("Server err: %v", err))
	}
}
