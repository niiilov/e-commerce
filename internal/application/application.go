package application

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/niiilov/e-commerce/cmd/docs"
	"github.com/niiilov/e-commerce/internal/user/del"
	"github.com/niiilov/e-commerce/internal/user/edit"
	"github.com/niiilov/e-commerce/internal/user/inf"
	"github.com/niiilov/e-commerce/internal/user/login"
	"github.com/niiilov/e-commerce/internal/user/reg"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

type APIs struct {
	router *mux.Router
	logger *zap.Logger
}

func New() *APIs {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	return &APIs{
		router: mux.NewRouter(),
		logger: logger,
	}
}

func (a *APIs) Routing() {

	// @swagger
	a.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	//user
	a.router.HandleFunc("/user/reg", reg.RegHandler).Methods("POST")         // registration
	a.router.HandleFunc("/user/login", login.LoginHandler).Methods("POST")   // authentification
	a.router.HandleFunc("/user/logout", login.LogoutHandler).Methods("POST") // exit
	a.router.HandleFunc("/user/inf", inf.InfHandler).Methods("GET")          // information output
	a.router.HandleFunc("/user/edit", edit.EditHandler).Methods("PUT")       // editing profile
	a.router.HandleFunc("/user/delete", del.DeleteHandler).Methods("DELETE") // delete profile

	// TODO: add handlers products, basket, orders and etc.
}

func (a *APIs) Run() error {
	a.logger.Info("Server starting on host :8080")
	return http.ListenAndServe(":8080", a.router)
}
