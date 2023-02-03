package main

import (
	"context"
	"fmt"
	handle "mysrtafes-backend/http-handle"
	"mysrtafes-backend/pkg/challenge"
	"mysrtafes-backend/repository"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Env struct {
	Addr     string
	DBConfig DBConfig
}

var env = Env{
	Addr: os.Getenv("ADDR"),
	DBConfig: DBConfig{
		User: os.Getenv("MYS_RTA_FES_DB_USER"),
		Pass: os.Getenv("MYS_RTA_FES_DB_PASS"),
		Host: os.Getenv("MYS_RTA_FES_DB_HOST"),
		Port: os.Getenv("MYS_RTA_FES_DB_PORT"),
		Name: os.Getenv("MYS_RTA_FES_DB_NAME"),
	},
}

func init() {
	if env.Addr == "" {
		env.Addr = ":80"
	}
}

func main() {
	// DBの生成
	db, err := newDB(env.DBConfig)
	if err != nil {
		panic(err)
	}
	dbRepository := repository.New(db)
	// Serviceの生成
	services := handle.NewServices(
		env.Addr,
		challenge.NewServer(dbRepository),
	)

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
