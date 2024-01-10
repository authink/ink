package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/authink/ink.go/src/core"
)

func setupGracefulShutdown(ink *core.Ink, srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// 等待中断/运行终止信号，优雅地关闭服务器
	<-quit
	log.Println("Shutting down ...")

	// 设置 5 秒的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(ink.Env.ShutdownTimeout)*time.Second)
	defer cancel()

	// 在超时时间内停止服务
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown: %s\n", err)
	}
}
