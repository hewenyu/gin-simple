package main

import (
	"fmt"
	"log"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/hewenyu/gin-simple/initialize"
)

func main() {
	Router := initialize.Routers()

	endless.DefaultReadTimeOut = 10 * time.Millisecond
	endless.DefaultWriteTimeOut = 10 * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20

	endPoint := fmt.Sprintf(":%d", 8080)

	server := endless.NewServer(endPoint, Router)

	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
