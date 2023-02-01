package main

import (
	"fmt"
	"mysrtafes-backend/pkg/game"
	"time"
)

func main() {
	// TODO: APIの環境構築ロジックを書く
	g := game.New(
		1,
		"風来のシレン",
		"ローグライクの大人気作品",
		"チュンソフト",
		"チュンソフト",
		game.ReleaseDate(time.Date(1995, 12, 1, 0, 0, 0, 0, time.UTC)),
		nil,
		nil,
	)
	fmt.Println(g)
}
