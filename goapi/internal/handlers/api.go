package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/sks123456/goapi/internal/middleware"
)

func Handler(r *chi.Mux){
	//global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/accounts",func(router chi.Router){

		//Middleware for /account route
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}