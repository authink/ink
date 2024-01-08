package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	env := loadEnv()

	db := connectDB(env)

	ink := &Ink{
		env,
		db,
	}

	r := gin.Default()
	setupRouter(ink, r)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", env.Host, env.Port),
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// 等待中断/运行终止信号，优雅地关闭服务器
	<-quit
	log.Println("Shutting down ...")

	// 设置 5 秒的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(env.ShutdownTimeout)*time.Second)
	defer cancel()

	// 在超时时间内停止服务
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown: %s\n", err)
	}
}
