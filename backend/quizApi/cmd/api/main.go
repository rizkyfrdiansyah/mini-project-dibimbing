package main

import (
	"log"
	"quizApi/pkg/api"
)

func main() {
	app := api.Default()
	err := app.Start()
	if err != nil {
		log.Println(err)
		panic(err)
	}
}