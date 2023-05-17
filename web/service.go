package web

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func Service(address, port string) {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	Route(g)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", address, port),
		Handler: g,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 优雅重启或停机 https://gin-gonic.com/zh-cn/docs/examples/graceful-restart-or-stop/

	// 监听进程退出信号
	quitCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	//监听退出信号
	<-quitCtx.Done()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
