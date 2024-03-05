package main

import (
	"fmt"
	"prisoners-dilemma/internal/app"
	"prisoners-dilemma/internal/config"
)

func main() {
	conf := config.MustLoad()

	app := app.NewApp(conf)
	err := app.Run()

	if err != nil {
		fmt.Println(err)
	}
}
