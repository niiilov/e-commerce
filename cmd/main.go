package main

import (
	"github.com/niiilov/e-commerce/internal/application"
	"github.com/niiilov/e-commerce/internal/db"
)

func main() {
	db.InitDB()
	api := application.New()
	api.Routing()
	api.Run()
}
