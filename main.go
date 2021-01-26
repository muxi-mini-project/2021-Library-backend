package main

import (
	"Library/model"
)

func main() {
	DB := model.Initdb()
	defer DB.Close()
	router.Run()
}
