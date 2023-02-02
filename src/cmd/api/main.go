package main

import (
	"context"
	"fmt"
	handle "mysrtafes-backend/http-handle"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// TODO: envの設定
	services := handle.NewServices(":3000")

	// 終了シグナル受け取りContextの定義
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt, os.Kill)

	// APIサーバー起動
	server := services.Server()

	go func() {
		// 中断処理実行後に動作
		<-ctx.Done()
		// タイムアウトのcontextを生成
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// Graceful ShutDown
		server.Shutdown(ctx)
	}()

	fmt.Println("starting api server")
	server.ListenAndServe()
	fmt.Println("shutdown api server")
}
