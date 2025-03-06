package main

import "github.com/niiilov/e-commerce/internal/application"

func main() {
	api := application.New()
	api.Routing()
	api.Run()
}
