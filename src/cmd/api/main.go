package main

import (
	"context"
	"fmt"
	handle "mysrtafes-backend/handle/http"
	"mysrtafes-backend/pkg/challenge"
	"mysrtafes-backend/pkg/game"
	"mysrtafes-backend/pkg/game/platform"
	"mysrtafes-backend/pkg/game/tag"
	"mysrtafes-backend/repository"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 環境変数
type osEnv struct {
	Env      Env
	Addr     string
	DBConfig DBConfig
}

var env = osEnv{
	Env:  os.Getenv("MYS_RTA_FES_ENV"),
	Addr: os.Getenv("ADDR"),
	DBConfig: DBConfig{
		User: os.Getenv("MYS_RTA_FES_DB_USER"),
		Pass: os.Getenv("MYS_RTA_FES_DB_PASS"),
		Host: os.Getenv("MYS_RTA_FES_DB_HOST"),
		Port: os.Getenv("MYS_RTA_FES_DB_PORT"),
		Name: os.Getenv("MYS_RTA_FES_DB_NAME"),
	},
}

// 動作環境
type Env = string

const (
	Env_Dev        Env = "Dev"
	Env_Production Env = "Production"
)

// 初期化
func init() {
	if env.Addr == "" {
		env.Addr = ":80"
	}
	if env.Env == "" {
		env.Env = Env_Dev
	}
}

func main() {
	// DBの生成
	db, err := newDB(env.DBConfig, env.Env)
	if err != nil {
		panic(err)
	}
	dbRepository := repository.New(db)
	// Serviceの生成
	services := handle.NewServices(
		env.Addr,
		game.NewServer(dbRepository),
		challenge.NewServer(dbRepository),
		tag.NewServer(dbRepository),
		platform.NewServer(dbRepository),
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
